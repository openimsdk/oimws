package sdk

import (
	"context"
	"fmt"
	"github.com/openim-sigs/oimws/pkg/common/config"
	"github.com/openim-sigs/oimws/pkg/core_func"
	module2 "github.com/openim-sigs/oimws/pkg/module"
	"github.com/openimsdk/tools/utils/datautil"
	log "github.com/xuexihuang/new_log15"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(ctx context.Context, index int, conf *Config, logConf *config.Log) error {
	listenPort, err := datautil.GetElemByIndex(conf.SdkConfig.SdkWsPort, index)
	if err != nil {
		return err
	}
	core_func.Config.WsAddr = conf.SdkConfig.OpenimWs
	core_func.Config.ApiAddr = conf.SdkConfig.OpenimApi
	core_func.Config.DataDir = conf.SdkConfig.DbDir
	core_func.Config.LogLevel = uint32(logConf.RemainLogLevel)
	core_func.Config.IsLogStandardOutput = logConf.IsStdout
	core_func.Config.LogFilePath = logConf.StorageLocation
	core_func.Config.IsExternalExtensions = true
	log.SetOutLevel(log.LvlInfo)
	fmt.Println("Client starting....")
	log.Info("Client starting....")
	gatenet := Initsever(listenPort)
	gatenet.SetMsgFun(module2.NewAgent, module2.CloseAgent, module2.DataRecv)
	go gatenet.Runloop()
	module2.ProgressStartTime = time.Now().Unix()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGTERM)
	sig := <-c
	log.Info("wsconn server closing down ", "sig", sig)
	gatenet.CloseGate()
	return nil
}
