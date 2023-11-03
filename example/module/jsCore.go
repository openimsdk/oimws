package module

import (
	"encoding/json"
	"fmt"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/utils"
	"reflect"

	"github.com/xuexihuang/new_gonet/example/core_func"
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
	//recv消息
	RecvMsg() chan interface{} //todo your sturct,error or response
	//send消息
	SendMsg(interface{}) error
	//关闭循环，并释放资源
	Destroy()
}

func NewJsCore(para *ParamStru) *JsCore {
	respChan := make(chan *core_func.EventData)
	funcRouter := core_func.NewFuncRouter(respChan)
	funcRouter.InitSDK(para.GetOperationID(), para.GetPlatformID())
	return &JsCore{RespMessagesChan: respChan, funcRouter: funcRouter}
}
func (core *JsCore) RecvMsg() chan *core_func.EventData {
	return core.RespMessagesChan
}

func (core *JsCore) SendMsg(req *Req) error {
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
		argsValue[i] = reflect.ValueOf(arg)
	}
	methodValue.Call(argsValue)
	return nil
}
func (core *JsCore) Destroy() {
	core.funcRouter.Logout("socket close")
}
