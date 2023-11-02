package module

import (
	"encoding/json"
	"fmt"
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
}

type JsInterface interface {
	//recv消息
	RecvMsg() chan interface{} //todo your sturct,error or response
	//send消息
	SendMsg(interface{}) error
	//关闭循环，并释放资源
	Destroy()
}
type jsParam struct {
	userId       string
	token        string
	platformID   string
	operationID  string
	isBackground bool
}

func NewJsCore(para *ParamStru) *JsCore {
	respChan := make(chan *core_func.EventData)
	return &JsCore{RespMessagesChan: respChan, funcRouter: core_func.NewFuncRouter(respChan)}
}
func (core *JsCore) RecvMsg() chan *core_func.EventData {

	return core.RespMessagesChan
}

func (core *JsCore) SendMsg(req *Req) error {
	methodValue := reflect.ValueOf(core.funcRouter).MethodByName(req.ReqFuncName)
	if !methodValue.IsValid() {
		//log.ZWarn(context.Background(), "method is valid", errors.New("method is valid"), "data", req)
		fmt.Println("method is valid", "data=", req)
		//todo return err info with operationID
	}
	var args []any
	if err := json.Unmarshal([]byte(req.Data), &args); err != nil {
		//todo todo return err info with operationID
	}
	// Convert args to []reflect.Value
	argsValue := make([]reflect.Value, len(args))
	for i, arg := range args {
		argsValue[i] = reflect.ValueOf(arg)
	}
	methodValue.Call(argsValue)
	//core.ReceivMsgChan<-req
	return nil
}
func (core *JsCore) Destroy() {

	// destroy funcRouter
}
