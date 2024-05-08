package sdk

import (
	"context"
	"fmt"
	"github.com/openim-sigs/oimws/pkg/common/config"
	"github.com/openim-sigs/oimws/pkg/core_func"
	"github.com/openim-sigs/oimws/pkg/module"
	log "github.com/xuexihuang/new_log15"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(ctx context.Context, index int, conf *Config, logConf *config.Log) error {
	listenPort, err := datautilGetElemByIndex(conf.SdkConfig.SdkWsPort, index)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(conf.SdkConfig.DbDir, os.ModePerm); err != nil && !os.IsExist(err) {
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
	gatenet.SetMsgFun(module.NewAgent, module.CloseAgent, module.DataRecv)
	go gatenet.Runloop()
	module.ProgressStartTime = time.Now().Unix()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGTERM)
	sig := <-c
	log.Info("wsconn server closing down ", "sig", sig)
	gatenet.CloseGate()
	return nil
}

// TODO: datautil.GetElemByIndex
func datautilGetElemByIndex(array []int, index int) (int, error) {
	if index < 0 || index >= len(array) {
		return 0, fmt.Errorf("index out of range index %d array %+v", index, array)
	}
	return array[index], nil
}
