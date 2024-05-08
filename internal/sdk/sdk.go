package sdk

import (
	"fmt"
	"github.com/openim-sigs/oimws/pkg/common/config"
	gate2 "github.com/openim-sigs/oimws/pkg/gate"
	"github.com/openim-sigs/oimws/pkg/network/tjson"
	"sync"
	"time"
)

const (
	MaxMsgLen     = 1024 * 1024 * 10
	MaxConnNum    = 100 * 100 * 10
	HTTPTimeout   = 10 * time.Second
	WriterChanLen = 1000
)

type Config struct {
	SdkConfig config.Config
}

var Processor = tjson.NewProcessor()

type GateNet struct {
	*gate2.Gate
	CloseSig chan bool
	Wg       sync.WaitGroup
}

// Initsever initializes a new GateNet instance with the given WebSocket port and default configurations.
func Initsever(wsPort int) *GateNet {
	gatenet := new(GateNet)
	gatenet.Gate = gate2.NewGate(MaxConnNum, MaxMsgLen,
		Processor, fmt.Sprintf(":%d", wsPort), HTTPTimeout, WriterChanLen)
	gatenet.CloseSig = make(chan bool, 1)
	return gatenet
}

// SetMsgFun sets the functions that will be called on new connection, disconnection, and data reception events.
func (gt *GateNet) SetMsgFun(Fun1 func(gate2.Agent), Fun2 func(gate2.Agent), Fun3 func(interface{}, gate2.Agent)) {
	gt.Gate.SetFun(Fun1, Fun2, Fun3)
}

// Runloop starts the server loop in a new goroutine and ensures the WaitGroup is properly managed.
func (gt *GateNet) Runloop() {
	gt.Wg.Add(1)
	gt.Run(gt.CloseSig)
	gt.Wg.Done()
}

// CloseGate sends a signal to close the server and waits for all goroutines to finish before calling OnDestroy.
func (gt *GateNet) CloseGate() {
	gt.CloseSig <- true
	gt.Wg.Wait()
	gt.Gate.OnDestroy()
}
