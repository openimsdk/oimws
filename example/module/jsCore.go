package module

type JsCore struct {
}

type JsInterface interface {
	//recv消息
	RecvMsg() chan interface{} //todo your sturct,error or response
	//send消息
	SendMsg(interface{}) error
	//关闭循环，并释放资源
	Destroy()
	//
	run()
}

func NewJsCore() *JsCore {
	return nil
}
func (core *JsCore) RecvMsg() chan interface{} {
	return nil
}
func (core *JsCore) SendMsg(data interface{}) error {

	return nil
}
func (core *JsCore) Destroy() {

}
func (core *JsCore) run() {

}
