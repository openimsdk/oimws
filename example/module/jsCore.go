package module

type JsCore struct {
	closeChan       chan bool
	ReceivMsgChan   chan interface{}
	OutMsgChan      chan interface{}
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
	return nil
}
func (core *JsCore) RecvMsg() chan interface{} {

	return core.OutMsgChan
}

func (core *JsCore) SendMsg(data interface{}) error {

	core.ReceivMsgChan<-data
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
		case indata:=<-core.ReceivMsgChan
		//todo your logic
		case <sdk.out
		}
	}

}
