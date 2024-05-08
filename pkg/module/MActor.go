package module

import (
	"context"
	"encoding/json"
	"errors"
	common2 "github.com/openim-sigs/oimws/pkg/common"
	"github.com/openim-sigs/oimws/pkg/core_func"
	"github.com/openim-sigs/oimws/pkg/gate"
	"net/url"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/xuexihuang/new_log15"
)

const (
	WsUserID    = "sendID"
	OperationID = "operationID"
	PlatformID  = "platformID"
)
const ProtocolError = "Protocol Error"
const DisconnectGCLimit = 100

var disConnectNum atomic.Int64

type ParamStru struct {
	UrlPath   string
	Token     string
	SessionId string
	UserId    int64
	GroupId   int64
	OrgId     int64
	OrgName   string
}

// GetUserID parses the URL to get the UserID parameter.
func (p *ParamStru) GetUserID() string {
	u, err := url.Parse(p.UrlPath)
	if err != nil {
		return ""
	}
	return u.Query().Get(WsUserID)
}

// GetOperationID parses the URL to get the OperationID parameter.
func (p *ParamStru) GetOperationID() string {
	u, err := url.Parse(p.UrlPath)
	if err != nil {
		return ""
	}
	return u.Query().Get(OperationID)
}

// GetPlatformID parses the URL to get the PlatformID parameter.
func (p *ParamStru) GetPlatformID() string {
	u, err := url.Parse(p.UrlPath)
	if err != nil {
		return ""
	}
	return u.Query().Get(PlatformID)
}

type ResReleaseStru struct {
	BackSign chan bool
}
type MActorIm struct {
	//todo your module ojb values
	mJsCore          *JsCore
	heartTickerSend  *time.Ticker //用于心跳send
	param            *ParamStru
	nChanLen         int //接收数据网络缓存
	wg               sync.WaitGroup
	a                gate.Agent
	SessionId        string
	closeChan        chan bool //主动关闭协程的通道
	releaseResChan   chan *ResReleaseStru
	ReceivMsgChan    chan interface{} //接收网络层数据通道
	heartTicker      *time.Ticker     //用于心跳监测
	heartFlag        bool             //初始为false，收到心跳pack设置为true
	isclosing        bool
	isReleasedJscore bool
}

// NewMActor creates a new actor instance.
func NewMActor(a gate.Agent, sessionId string, appParam *ParamStru) (MActor, error) {
	ret := &MActorIm{param: appParam, a: a, SessionId: sessionId, releaseResChan: make(chan *ResReleaseStru, 1), closeChan: make(chan bool, 1), nChanLen: 10, ReceivMsgChan: make(chan interface{}, 10), isclosing: false,
		heartTicker: time.NewTicker(100 * time.Second), heartFlag: false, heartTickerSend: time.NewTicker(28 * time.Second), isReleasedJscore: false}
	///////////////////////////////////////
	ret.mJsCore = NewJsCore(appParam, sessionId) //todo
	///////////////////////////////////////
	go ret.run()
	return ret, nil
}

// run contains the main loop for the actor, handling various operations.
func (actor *MActorIm) run() {
	actor.wg.Add(1)
	defer common2.TryRecoverAndDebugPrint()
	defer actor.wg.Done()
	for {
		select {
		case <-actor.heartTickerSend.C: //send the heart pack
			if actor.isclosing == true {
				continue
			}
			actor.sendHeart()
		case <-actor.closeChan:
			log.Info("收到退出信号", "sessionId", actor.SessionId)
			if !actor.isReleasedJscore {
				actor.mJsCore.Destroy()
			}
			if disConnectNum.Add(1) > DisconnectGCLimit {
				runtime.GC()
				disConnectNum.Store(0)
			}
			return
		case resChan := <-actor.releaseResChan:
			log.Info("收到释放资源通道消息")
			actor.mJsCore.Destroy()
			actor.a.Destroy()
			actor.isReleasedJscore = true
			resChan.BackSign <- true
		case recvData := <-actor.ReceivMsgChan:
			if actor.isclosing == true {
				continue
			}
			data := recvData.(*common2.TWSData)
			_ = actor.doRecvPro(data)
		case resp := <-actor.mJsCore.RecvMsg():
			actor.sendEventResp(resp)
			if resp.Event == LogoutName {
				actor.isReleasedJscore = true
				actor.isclosing = true
				actor.sendClosingResp()
			}
			//case <-actor.heartTicker.C:
			//	if actor.heartFlag == true {
			//		actor.heartFlag = false
			//	} else {
			//		log.Error("心跳包超时错误", "sessionId", actor.SessionId)
			//		actor.isclosing = true
			//		actor.a.Destroy()
			//	}
		}
	}
}
func (actor *MActorIm) ReleaseRes() {
	log.Info("get ReleaseRes sign")
	ind := &ResReleaseStru{BackSign: make(chan bool, 1)}
	actor.releaseResChan <- ind
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		log.Info("出现超时返回，actor可能已经被异步destroy")
	case <-ind.BackSign:
		log.Info("通过releaseRes接口回收资源")
	}

}
func (actor *MActorIm) Destroy() {
	actor.closeChan <- true
	actor.wg.Wait()
	actor.a = nil
	log.Info("退出MQPushActorIm", "sessionId", actor.SessionId)
}

// ProcessRecvMsg processes received messages and sends them to the ReceivMsgChan.
func (actor *MActorIm) ProcessRecvMsg(msg interface{}) error {
	if len(actor.ReceivMsgChan) == actor.nChanLen {
		log.Error("send channel is full", "sessionId", actor.SessionId)
		return errors.New("send channel is full")
	}
	actor.ReceivMsgChan <- msg
	return nil
}

// doRecvPro processes the message received from the network layer.
func (actor *MActorIm) doRecvPro(data *common2.TWSData) error {
	log.Info("message come here", "data", data)
	if data.MsgType == common2.MessageText {
		req := &Req{}
		err := json.Unmarshal(data.Msg, req)
		if err != nil {
			log.Error("parse protocol err", "err", err, "sessionId", actor.SessionId)
			actor.sendEventResp(&core_func.EventData{Event: ProtocolError, ErrCode: 20000, ErrMsg: err.Error(),
				OperationID: req.OperationID})
			return err
		}
		log.Info("receive req", "req", req, "sessionId", actor.SessionId)
		err = actor.mJsCore.SendMsg(req)
		if err != nil {
			actor.sendEventResp(&core_func.EventData{Event: req.ReqFuncName, ErrCode: 20000, ErrMsg: err.Error(),
				OperationID: req.OperationID})
		}
	}
	return nil
}

// sendResp sends a response message to the WebSocket client.
func (actor *MActorIm) sendHeart() {
	//heart := []byte("ping")
	resSend := &common2.TWSData{MsgType: common2.PingMessage, Msg: nil}
	actor.a.WriteMsg(resSend)
}

// sendEventResp sends an event response to the WebSocket client.
func (actor *MActorIm) sendEventResp(res *core_func.EventData) {
	resb, _ := json.Marshal(res)
	resSend := &common2.TWSData{MsgType: common2.MessageText, Msg: resb}
	actor.a.WriteMsg(resSend)
}

func (actor *MActorIm) sendClosingResp() {
	resSend := &common2.TWSData{MsgType: common2.CloseMessage, Msg: nil}
	actor.a.WriteMsg(resSend)
}
