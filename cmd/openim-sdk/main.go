package main

import (
	"flag"
	"fmt"
	"github.com/openim-sigs/oimws/pkg/core_func"
	gate2 "github.com/openim-sigs/oimws/pkg/gate"
	module2 "github.com/openim-sigs/oimws/pkg/module"
	"github.com/openim-sigs/oimws/pkg/network/tjson"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/xuexihuang/new_log15"
)

const (
	MaxMsgLen     = 1024 * 1024 * 10
	MaxConnNum    = 100 * 100 * 10
	HTTPTimeout   = 10 * time.Second
	WriterChanLen = 1000
)

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
		Processor, ":"+fmt.Sprintf("%d", wsPort), HTTPTimeout, WriterChanLen)
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

// The main function sets up the WebSocket server and handles graceful shutdowns.
func main() {
	var sdkWsPort, logLevel *int
	var openIMWsAddress, openIMApiAddress, openIMDbDir *string
	openIMApiAddress = flag.String("openIM_api_address", "http://127.0.0.1:10002",
		"openIM api listening address")
	openIMWsAddress = flag.String("openIM_ws_address", "ws://127.0.0.1:10001",
		"openIM ws listening address")
	sdkWsPort = flag.Int("sdk_ws_port", 10003, "openIMSDK ws listening port")
	logLevel = flag.Int("openIM_log_level", 5, "control log output level")
	openIMDbDir = flag.String("openIMDbDir", "./db", "openIM db dir")
	flag.Parse()
	core_func.Config.WsAddr = *openIMWsAddress
	core_func.Config.ApiAddr = *openIMApiAddress
	core_func.Config.DataDir = *openIMDbDir
	core_func.Config.LogLevel = uint32(*logLevel)
	core_func.Config.IsLogStandardOutput = true
	log.SetOutLevel(log.LvlInfo)
	fmt.Println("Client starting....")
	log.Info("Client starting....")
	gatenet := Initsever(*sdkWsPort)
	gatenet.SetMsgFun(module2.NewAgent, module2.CloseAgent, module2.DataRecv)
	go gatenet.Runloop()
	/////////////////////////////////////
	//statusGate := Initsever(90)
	//gatenet.SetMsgFun(module.NewStatusAgent, module.CloseStatusAgent, module.DataRecvStatus)
	//go gatenet.Runloop()
	module2.ProgressStartTime = time.Now().Unix()
	///////////////////////////////////////////
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGTERM)
	sig := <-c
	log.Info("wsconn server closing down ", "sig", sig)
	gatenet.CloseGate()
	//statusGate.CloseGate()
}
