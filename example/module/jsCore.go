package module

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OpenIMSDK/tools/log"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/sdkerrs"
	"github.com/xuexihuang/new_gonet/example/core_func"
	"reflect"
)
/*
	JsCore 相关的内容，主要实现了信息收、发、销毁
*/

type JsCore struct {
	closeChan       chan bool
	ReceivMsgChan   chan interface{}
	OutMsgChan      chan interface{}
	respMessagesChan chan *core_func.EventData
	funcRouter      *core_func.FuncRouter
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


func NewJsCore() *JsCore {
	return &JsCore{funcRouter: core_func.NewFuncRouter()}
}
func (core *JsCore) RecvMsg() chan interface{} {
	return core.OutMsgChan
}

func (core *JsCore) SendMsg(request interface{}) error{
//func (core *JsCore) SendMsg(req *Req) error {
	req:=request.(Req)
	methodValue := reflect.ValueOf(core.funcRouter).MethodByName(req.ReqFuncName)
	if !methodValue.IsValid() {
		log.ZWarn(context.Background(),"method is valid",errors.New("method is valid"),
			"data",req)
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
	// TODO
	methodValue.Call(argsValue)
	//core.ReceivMsgChan<-req
	return nil
}
func (core *JsCore) Destroy() {

	core.closeChan<-true
}

func (core *JsCore) run() {
	for  {
		select {
		case <-core.closeChan:
			return
		case indata:=<-core.ReceivMsgChan:
		//todo your logic
		case <-sdk.out:
		}
	}

}
、