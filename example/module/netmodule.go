package module

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/openim-sigs/oimws/common"
	"github.com/openim-sigs/oimws/gate"
	log "github.com/xuexihuang/new_log15"
)

type MActor interface {
	//recv消息
	ProcessRecvMsg(interface{}) error
	//关闭循环，并释放资源
	Destroy()
	//
	run()
}

func NewAgent(a gate.Agent) {
	aUerData := a.UserData().(*common.TAgentUserData)
	log.Info("one ws connect", "sessionId", aUerData.SessionID)
	param, err := checkToken(aUerData)
	if err != nil {
		log.Error("token校验失败", "userData", aUerData, "sessionId", aUerData.SessionID)
		res := &ResponseSt{Type: RESP_OP_TYPE, Cmd: CONN_CMD, Success: false, ErrMsg: "check token error"}
		resb, _ := json.Marshal(res)
		resSend := &common.TWSData{MsgType: common.MessageText, Msg: resb}
		a.WriteMsg(resSend)
		a.Close()
		return
	}
	log.Info("checkToken info", "param", param, "err", err)
	actor, err := NewMActor(a, param.SessionId, param)
	if err != nil {
		log.Error("NewMQActor error", "err", err, "sessionId", aUerData.SessionID)
		res := &ResponseSt{Type: RESP_OP_TYPE, Cmd: CONN_CMD, Success: false, ErrMsg: "NewMQActor error"}
		resb, _ := json.Marshal(res)
		resSend := &common.TWSData{MsgType: common.MessageText, Msg: resb}
		a.WriteMsg(resSend)
		a.Close()
		return
	}
	aUerData.ProxyBody = actor
	a.SetUserData(aUerData)
	log.Info("one linked", "param", param, "sessionId", aUerData.SessionID)
}

func CloseAgent(a gate.Agent) {
	aUerData := a.UserData().(*common.TAgentUserData)
	if aUerData.ProxyBody != nil {
		aUerData.ProxyBody.(MActor).Destroy()
		aUerData.ProxyBody = nil
	}
	log.Info("one dislinkder", "sessionId", a.UserData().(*common.TAgentUserData).SessionID)
}
func DataRecv(data interface{}, a gate.Agent) {
	aUerData := a.UserData().(*common.TAgentUserData)
	if aUerData.ProxyBody != nil {
		err := aUerData.ProxyBody.(MActor).ProcessRecvMsg(data)
		if err != nil {
			log.Error("溢出错误", "sessionId", aUerData.SessionID)
			a.Destroy()
		}
	}
}
func checkToken(data *common.TAgentUserData) (*ParamStru, error) {
	ret := new(ParamStru)
	ret.SessionId = data.SessionID
	var token string
	if data.CookieVal != "" {
		token = data.CookieVal
	} else {
		/////////////////////
		u, err := url.Parse(data.AppString)
		if err != nil {
			log.Error("ws url path not correct", "sessionId", data.SessionID)
			return nil, errors.New("ws url path not correct")
		}
		q := u.Query()
		token = q.Get("token")
		//////////////////////
	}
	if token == "" {
		log.Error("获取token为空", "sessionId", data.SessionID)
		return nil, errors.New("获取token为空")
	}
	//todo  这里添加你的token效验逻辑验证token的合法性
	//ret.UserId=""
	ret.UrlPath = data.AppString
	ret.Token = token
	return ret, nil
}
