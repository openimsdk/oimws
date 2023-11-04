package module

import (
	"encoding/json"
	"errors"
	"net/url"
	"sync"
	"time"

	"github.com/openim-sigs/oimws/example/core_func"

	"github.com/openim-sigs/oimws/common"
	"github.com/openim-sigs/oimws/gate"
	log "github.com/xuexihuang/new_log15"
)

const (
	WsUserID    = "sendID"
	OperationID = "operationID"
	PlatformID  = "platformID"
)

type ParamStru struct {
	UrlPath   string
	Token     string
	SessionId string
	UserId    int64
	GroupId   int64
	OrgId     int64
	OrgName   string
}

// GetOperationID parses the URL to get the OperationID parameter
func (p *ParamStru) GetOperationID() string {
	u, err := url.Parse(p.UrlPath)
	if err != nil {
		return ""
	}
	return u.Query().Get(OperationID)
}

// GetPlatformID parses the URL to get the PlatformID parameter
func (p *ParamStru) GetPlatformID() string {
	u, err := url.Parse(p.UrlPath)
	if err != nil {
		return ""
	}
	return u.Query().Get(PlatformID)
}

type MActorIm struct {
	//todo your module ojb values
	mJsCore         *JsCore
	heartTickerSend *time.Ticker //用于心跳send
	param           *ParamStru
	nChanLen        int //接收数据网络缓存
	wg              sync.WaitGroup
	a               gate.Agent
	SessionId       string
	closeChan       chan bool        //主动关闭协程的通道
	ReceivMsgChan   chan interface{} //接收网络层数据通道
	heartTicker     *time.Ticker     //用于心跳监测
	heartFlag       bool             //初始为false，收到心跳pack设置为true
	isclosing       bool
}

// NewMActor creates a new actor instance
func NewMActor(a gate.Agent, sessionId string, appParam *ParamStru) (MActor, error) {
	ret := &MActorIm{param: appParam, a: a, SessionId: sessionId, closeChan: make(chan bool, 1), nChanLen: 10, ReceivMsgChan: make(chan interface{}, 10), isclosing: false,
		heartTicker: time.NewTicker(100 * time.Second), heartFlag: false, heartTickerSend: time.NewTicker(100 * time.Second)}
	///////////////////////////////////////
	ret.mJsCore = NewJsCore(appParam, sessionId) //todo
	///////////////////////////////////////
	go ret.run()
	return ret, nil
}

// run contains the main loop for the actor, handling various operations
func (actor *MActorIm) run() {
	actor.wg.Add(1)
	defer common.TryRecoverAndDebugPrint()
	defer actor.wg.Done()
	for {
		select {
		case <-actor.heartTickerSend.C: //send the heart pack
			if actor.isclosing == true {
				continue
			}
			actor.sendResp(&ResponseSt{Type: "heart"})
		case <-actor.closeChan:
			log.Info("收到退出信号", "sessionId", actor.SessionId)
			actor.mJsCore.Destroy()
			return
		case recvData := <-actor.ReceivMsgChan:
			if actor.isclosing == true {
				continue
			}
			data := recvData.(*common.TWSData)
			_ = actor.doRecvPro(data)
		case resp := <-actor.mJsCore.RecvMsg():
			//if jscoredata.ErrCode != 0 {
			//	actor.sendResp(nil) //todo send errormsg
			//	actor.isclosing = true
			//	actor.a.Destroy()
			//} else {
			actor.sendEventResp(resp) // todo send msg
			//}
		case <-actor.heartTicker.C:
			if actor.heartFlag == true {
				actor.heartFlag = false
			} else {
				log.Error("心跳包超时错误", "sessionId", actor.SessionId)
				actor.isclosing = true
				actor.a.Destroy()
			}
		}
	}
}

// Destroy gracefully shuts down the actor
func (actor *MActorIm) Destroy() {
	actor.closeChan <- true
	actor.wg.Wait()
	actor.a = nil
	log.Info("退出MQPushActorIm", "sessionId", actor.SessionId)
}

// ProcessRecvMsg processes received messages and sends them to the ReceivMsgChan
func (actor *MActorIm) ProcessRecvMsg(msg interface{}) error {
	if len(actor.ReceivMsgChan) == actor.nChanLen {
		log.Error("send channel is full", "sessionId", actor.SessionId)
		return errors.New("send channel is full")
	}
	actor.ReceivMsgChan <- msg
	return nil
}

// doRecvPro processes the message received from the network layer
func (actor *MActorIm) doRecvPro(data *common.TWSData) error {
	log.Info("message come here", "data", data)
	if data.MsgType == common.MessageBinary {
		req := &Req{}
		err := json.Unmarshal(data.Msg, req)
		if err != nil {
			log.Error("解析前端协议出错", "err", err, "sessionId", actor.SessionId)
			//todo response error
			return err
		}
		log.Info("收到sub命令", "req", req, "sessionId", actor.SessionId)
		err = actor.mJsCore.SendMsg(req)
		if err != nil {
			actor.sendEventResp(&core_func.EventData{ErrCode: 20000, ErrMsg: err.Error(), OperationID: req.OperationID})
		}
	}
	return nil
}

// sendResp sends a response message to the WebSocket client
func (actor *MActorIm) sendResp(res *ResponseSt) {
	//heart := []byte("ping")
	resSend := &common.TWSData{MsgType: common.PingMessage, Msg: nil}
	actor.a.WriteMsg(resSend)
}

// sendEventResp sends an event response to the WebSocket client
func (actor *MActorIm) sendEventResp(res *core_func.EventData) {
	resb, _ := json.Marshal(res)
	resSend := &common.TWSData{MsgType: common.MessageBinary, Msg: resb}
	actor.a.WriteMsg(resSend)
}
