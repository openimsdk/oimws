package core_func

import (
	"runtime"
	"strings"
)

type ConnCallback struct {
	respMessage *RespMessage
}

func NewConnCallback(respMessage *RespMessage) *ConnCallback {
	return &ConnCallback{respMessage: respMessage}
}

func (c ConnCallback) OnConnecting() {
	c.respMessage.sendEventSuccessRespNoData(getSelfFuncName())
}

func (c ConnCallback) OnConnectSuccess() {
	c.respMessage.sendEventSuccessRespNoData(getSelfFuncName())
}

func (c ConnCallback) OnConnectFailed(errCode int32, errMsg string) {
	c.respMessage.sendEventFailedRespNoData(getSelfFuncName(), errCode, errMsg)
}

func (c ConnCallback) OnKickedOffline() {
	c.respMessage.sendEventSuccessRespNoData(getSelfFuncName())
}

func (c ConnCallback) OnUserTokenExpired() {
	c.respMessage.sendEventSuccessRespNoData(getSelfFuncName())
}

func getSelfFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	end := strings.LastIndex(runtime.FuncForPC(pc).Name(), ".")
	if end == -1 {
		return ""
	}
	return runtime.FuncForPC(pc).Name()[end+1:]
}
