package main

import (
	"fmt"
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

func Initsever() *GateNet {
	gatenet := new(GateNet)
	gatenet.Gate = &gate.Gate{
		MaxConnNum:      100,
		PendingWriteNum: 200,
		MaxMsgLen:       20000,
		WSAddr:          ":80", // todo change your ip_port
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
	log.SetOutLevel(log.LvlInfo)
	fmt.Println("客户端启动....")
	log.Info("客户端启动....")
	gatenet := Initsever()
	gatenet.SetMsgFun(module.NewAgent, module.CloseAgent, module.DataRecv)
	go gatenet.Runloop()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGTERM)
	sig := <-c
	log.Info("wsconn server closing down ", "sig", sig)
	gatenet.CloseGate()
}
