package core_func

import (
	"fmt"
	"testing"
)

var fu FuncRouter

func TestGetAllConversationList(T *testing.T) {

	ev := make(chan *EventData, 10)
	fu := NewFuncRouter(ev, "123456")
	fu.GetAllConversationList("11111")

	msg := <-fu.respMessage.respMessagesChan
	fmt.Println("msg:", msg)
}

type eventData struct {
	Event       string `json:"event"`
	ErrCode     int32  `json:"errCode"`
	ErrMsg      string `json:"errMsg"`
	Data        string `json:"data"`
	OperationID string `json:"operationID"`
}
