package core_func

import (
	"fmt"
	"testing"
)

func TestGetAllConversationList(T *testing.T) {

	ev := make(chan *EventData, 10)
	fu := NewFuncRouter(ev, "123456")
	fu.GetAllConversationList("11111")

	msg := <-fu.respMessage.respMessagesChan
	fmt.Println("msg:", msg)
}
