package module

import (
	"encoding/json"
	"fmt"
	"github.com/openim-sigs/oimws/pkg/core_func"
	"reflect"

	"github.com/openimsdk/openim-sdk-core/v3/pkg/utils"
)

const (
	LogoutTips = "js sdk socket close"
	LogoutName = "Logout"
)

type JsCore struct {
	RespMessagesChan chan *core_func.EventData
	funcRouter       *core_func.FuncRouter
}

type Req struct {
	ReqFuncName string `json:"reqFuncName" `
	OperationID string `json:"operationID"`
	Data        string `json:"data"`
	UserID      string `json:"userID"`
	Batch       int    `json:"batchMsg"`
}
type JsInterface interface {
	RecvMsg() chan interface{} //todo your sturct,error or response
	SendMsg(interface{}) error
	Destroy()
}

// NewJsCore creates a new JsCore instance.
func NewJsCore(para *ParamStru, sessionId string) *JsCore {
	respChan := make(chan *core_func.EventData, 100)
	funcRouter := core_func.NewFuncRouter(respChan, sessionId)
	fmt.Println("NewJsCore", "data=", "sessionId", sessionId)
	funcRouter.InitSDK(para.GetOperationID(), para.GetPlatformID())
	return &JsCore{RespMessagesChan: respChan, funcRouter: funcRouter}
}

// RecvMsg returns the channel to receive messages.
func (core *JsCore) RecvMsg() chan *core_func.EventData {
	return core.RespMessagesChan
}

// SendMsg processes the incoming request and calls the corresponding method.
func (core *JsCore) SendMsg(req *Req) error {
	fmt.Println("method is valid", "data=", req)
	methodValue := reflect.ValueOf(core.funcRouter).MethodByName(req.ReqFuncName)
	if !methodValue.IsValid() {
		//log.ZWarn(context.Background(), "method is valid", errors.New("method is valid"), "data", req)
		fmt.Println("method is valid", "data=", req)
		return utils.Wrap(fmt.Errorf("method is valid"), "method is valid")
	}
	var args []any
	if err := json.Unmarshal([]byte(req.Data), &args); err != nil {
		return utils.Wrap(err, "json.Unmarshal failed")
	}
	// Convert args to []reflect.Value
	args = append([]any{req.OperationID}, args...)
	argsValue := make([]reflect.Value, len(args))
	for i, arg := range args {
		if arg == nil {
			return utils.Wrap(fmt.Errorf("args[%d] is not nil", i), "args has nil")
		}
		argsValue[i] = reflect.ValueOf(arg)
	}
	methodValue.Call(argsValue)
	return nil
}

// Destroy performs cleanup when the core is no longer needed.
func (core *JsCore) Destroy() {
	core.funcRouter.Logout(LogoutTips)
}
