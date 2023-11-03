package main

import (
	"flag"
	"fmt"
	"github.com/xuexihuang/new_gonet/example/core_func"
	"github.com/xuexihuang/new_gonet/example/module"
	"github.com/xuexihuang/new_gonet/gate"
	"github.com/xuexihuang/new_gonet/network/tjson"
	log "github.com/xuexihuang/new_log15"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var Processor = tjson.NewProcessor()

type GateNet struct {
	*gate.Gate
	CloseSig chan bool
	Wg       sync.WaitGroup
}

func Initsever(wsPort int) *GateNet {
	gatenet := new(GateNet)
	gatenet.Gate = &gate.Gate{
		MaxConnNum:      100,
		PendingWriteNum: 200,
		MaxMsgLen:       20000,
		WSAddr:          ":" + fmt.Sprintf("%d", wsPort),
		HTTPTimeout:     10 * time.Second,
		CertFile:        "",
		KeyFile:         "",
		LenMsgLen:       2,
		Processor:       Processor,
	}
	gatenet.CloseSig = make(chan bool, 1)
	return gatenet
}

func (gt *GateNet) SetMsgFun(Fun1 func(gate.Agent), Fun2 func(gate.Agent), Fun3 func(interface{}, gate.Agent)) {
	gt.Gate.SetFun(Fun1, Fun2, Fun3)
}
func (gt *GateNet) Runloop() {
	gt.Wg.Add(1)
	gt.Run(gt.CloseSig)
	gt.Wg.Done()
}
func (gt *GateNet) CloseGate() {
	gt.CloseSig <- true
	gt.Wg.Wait()
	gt.Gate.OnDestroy()
}

func main() {
	var sdkWsPort, logLevel *int
	var openIMWsAddress, openIMApiAddress, openIMDbDir *string
	openIMApiAddress = flag.String("openIM_api_address", "http://127.0.0.1:10002",
		"openIM api listening address")
	openIMWsAddress = flag.String("openIM_ws_address", "ws://127.0.0.1:10001",
		"openIM ws listening address")
	sdkWsPort = flag.Int("sdk_ws_port", 10003, "openIMSDK ws listening port")
	logLevel = flag.Int("openIM_log_level", 6, "control log output level")
	openIMDbDir = flag.String("openIMDbDir", "../db/sdk/", "openIM db dir")
	flag.Parse()
	core_func.Config.WsAddr = *openIMWsAddress
	core_func.Config.ApiAddr = *openIMApiAddress
	core_func.Config.DataDir = *openIMDbDir
	core_func.Config.LogLevel = uint32(*logLevel)
	log.SetOutLevel(log.LvlInfo)
	fmt.Println("客户端启动....")
	log.Info("客户端启动....")
	gatenet := Initsever(*sdkWsPort)
	gatenet.SetMsgFun(module.NewAgent, module.CloseAgent, module.DataRecv)
	go gatenet.Runloop()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGTERM)
	sig := <-c
	log.Info("wsconn server closing down ", "sig", sig)
	gatenet.CloseGate()
}
