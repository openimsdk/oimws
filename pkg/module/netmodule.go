package module

import (
	"encoding/json"
	"errors"
	"github.com/openim-sigs/oimws/pkg/common"
	"github.com/openim-sigs/oimws/pkg/gate"
	log "github.com/xuexihuang/new_log15"
	"net/url"
	"sync"
)

type JsActorMap struct {
	sync.Mutex
	uActors map[string]MActor
}

var GJsActors *JsActorMap

func init() {
	GJsActors = &JsActorMap{uActors: make(map[string]MActor)}
}

type MActor interface {
	ProcessRecvMsg(interface{}) error
	Destroy()
	//
	ReleaseRes()
	run()
}

// NewAgent is called when a new WebSocket connection is established. It initializes agent-related data and checks the token validity.
func NewAgent(a gate.Agent) {
	aUerData := a.UserData().(*common.TAgentUserData)
	log.Info("one ws connect", "sessionId", aUerData.SessionID)
	param, err := checkToken(aUerData)
	if err != nil {
		log.Error("Token validation failed", "userData", aUerData, "sessionId", aUerData.SessionID)

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
	GJsActors.Lock()
	v, ok := GJsActors.uActors[param.GetUserID()]
	if ok {
		v.ReleaseRes()
	}
	GJsActors.uActors[param.GetUserID()] = actor
	GJsActors.Unlock()
	aUerData.ProxyBody = actor
	aUerData.UserId = param.GetUserID()
	a.SetUserData(aUerData)
	log.Info("one linked", "param", param, "sessionId", aUerData.SessionID)
}

// CloseAgent is called when the WebSocket connection is closed. It performs cleanup actions for the agent.
func CloseAgent(a gate.Agent) {
	aUerData := a.UserData().(*common.TAgentUserData)
	if aUerData.ProxyBody != nil {
		aUerData.ProxyBody.(MActor).Destroy()
		aUerData.ProxyBody = nil
	}
	GJsActors.Lock()
	_, ok := GJsActors.uActors[aUerData.UserId]
	if ok {
		delete(GJsActors.uActors, aUerData.UserId)
	}
	GJsActors.Unlock()
	log.Info("one dislinkder", "sessionId", a.UserData().(*common.TAgentUserData).SessionID)
}

// DataRecv is called when new data is received on the WebSocket connection. It processes the incoming data through the actor.
func DataRecv(data interface{}, a gate.Agent) {
	aUerData := a.UserData().(*common.TAgentUserData)
	if aUerData.ProxyBody != nil {
		err := aUerData.ProxyBody.(MActor).ProcessRecvMsg(data)
		if err != nil {
			log.Error("Overflow error", "sessionId", aUerData.SessionID)
			a.Destroy()
		}
	}
}

// checkToken validates the session token contained in the user data.
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
		log.Error("Token retrieval is empty", "sessionId", data.SessionID)
		return nil, errors.New("Token retrieval is empty")
	}
	// TODO: Add your token validation logic here to verify the legitimacy of the token
	//ret.UserId=""
	ret.UrlPath = data.AppString
	ret.Token = token
	if ret.GetUserID() == "" {
		log.Error("userId is empty!")
		return nil, errors.New("userId is empty")
	}
	return ret, nil
}
