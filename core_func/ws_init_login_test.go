package core_func

import (
	"fmt"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

type TestServer struct {
	fu          *FuncRouter
	operationID string
	sessionID   string
	platformID  string
}

const (
	operationID string = "123456"
	sessionID   string = "111"
	platformID  string = "1"
)

func NewTestServer() *TestServer {
	ev := make(chan *EventData, 10)
	return &TestServer{
		operationID: operationID,
		sessionID:   sessionID,
		platformID:  platformID,
		fu:          NewFuncRouter(ev, sessionID),
	}
}

func TestInitSDK(t *testing.T) {
	te := NewTestServer()
	fn := func() bool {
		te.fu.InitSDK(te.operationID, te.platformID)
		msg, err := <-te.fu.respMessage.respMessagesChan

		ret := &EventData{
			OperationID: te.operationID,
			Event:       "InitSDK",
			Data:        "",
		}
		assert.Equal(t, true, err)
		assert.Equal(t, ret, msg)
		return true
	}
	err := quick.Check(fn, nil)
	assert.Nil(t, err)
}

func TestLogin(t *testing.T) {
	te := NewTestServer()
	fn := func() bool {
		te.fu.Login(te.sessionID, te.operationID)
		msg, err := <-te.fu.respMessage.respMessagesChan

		ret := &EventData{
			OperationID: te.operationID,
			Event:       "Login",
			Data:        "",
		}
		fmt.Printf("ret,msg:%v", msg)
		assert.Equal(t, true, err)
		assert.Equal(t, ret, msg)
		return true
	}
	err := quick.Check(fn, nil)
	assert.Nil(t, err)
}

//func (f *FuncRouter) Login(operationID string, args ...any) {
//	f.setAllListener()
//	fmt.Println(operationID, "Login")
//	f.call(operationID, f.userForSDK.Login, args...)
//}
