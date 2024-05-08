package module

import (
	"encoding/json"
	common2 "github.com/openim-sigs/oimws/pkg/common"
	"github.com/openim-sigs/oimws/pkg/gate"
	log "github.com/xuexihuang/new_log15"
	"sync"
	"time"
)

type StatusActorIm struct {
	nChanLen        int          //接收数据网络缓存
	heartTickerSend *time.Ticker //用于心跳send
	wg              sync.WaitGroup
	a               gate.Agent
	SessionId       string
	closeChan       chan bool        //主动关闭协程的通道
	ReceivMsgChan   chan interface{} //接收网络层数据通道
	isclosing       bool
}

func NewStatusActor(a gate.Agent, sessionId string, appParam *ParamStru) (MActor, error) {

	ret := &StatusActorIm{a: a, SessionId: sessionId, closeChan: make(chan bool, 1), nChanLen: 10, ReceivMsgChan: make(chan interface{}, 10), isclosing: false,
		heartTickerSend: time.NewTicker(100 * time.Second)}
	go ret.run()
	return ret, nil
}
func (actor *StatusActorIm) run() {
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
			return
		case recvData := <-actor.ReceivMsgChan:
			if actor.isclosing == true {
				continue
			}
			data := recvData.(*common2.TWSData)
			_ = actor.doRecvPro(data)
		}
	}
}
func (actor *StatusActorIm) ProcessRecvMsg(interface{}) error {

	return nil
}
func (actor *StatusActorIm) Destroy() {
	actor.closeChan <- true
	actor.wg.Wait()
	actor.a = nil
	log.Info("退出MQPushActorIm", "sessionId", actor.SessionId)
}
func (actor *StatusActorIm) ReleaseRes() {

}
func (actor *StatusActorIm) sendHeart() {
	//heart := []byte("ping")
	resSend := &common2.TWSData{MsgType: common2.PingMessage, Msg: nil}
	actor.a.WriteMsg(resSend)
}
func (actor *StatusActorIm) doRecvPro(data *common2.TWSData) error {
	log.Info("message come here", "data.type", data.MsgType)
	if data.MsgType == common2.MessageText {
		req := &RequestSt{}
		err := json.Unmarshal(data.Msg, req)
		if err != nil {
			log.Error("解析前端协议出错", "err", err, "sessionId", actor.SessionId)
			return err
		}
		log.Info("收到命令", "req", req, "sessionId", actor.SessionId)
		////////////////////////////////////////////////
		res := ResponseSt{Type: RESP_OP_TYPE, Success: true, UserId: genGroupUserIds(), Duration: time.Now().Unix() - ProgressStartTime}
		resb, _ := json.Marshal(res)
		resSend := &common2.TWSData{MsgType: common2.MessageText, Msg: resb}
		actor.a.WriteMsg(resSend)
	}
	return nil
}
func genGroupUserIds() []string {
	var ret []string
	GJsActors.Lock()
	defer GJsActors.Unlock()
	for k, _ := range GJsActors.uActors {
		ret = append(ret, k)
	}
	return ret
}
