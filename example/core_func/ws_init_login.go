package core_func

import (
	"github.com/OpenIMSDK/tools/log"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/sdkerrs"
	"github.com/openimsdk/openim-sdk-core/v3/sdk_struct"
	"strconv"
)

var Config sdk_struct.IMConfig

const (
	rotateCount  uint = 0
	rotationTime uint = 24
)

func (f *FuncRouter) InitSDK(operationID, platformID string) {
	callback := NewConnCallback(f.respMessage)
	j, err := strconv.ParseInt(platformID, 10, 64)
	if err != nil {
		f.respMessage.sendOnErrorResp(operationID, err)
		return
	}
	config := sdk_struct.IMConfig{
		PlatformID:           int32(j),
		ApiAddr:              Config.ApiAddr,
		WsAddr:               Config.WsAddr,
		DataDir:              Config.DataDir,
		LogLevel:             Config.LogLevel,
		IsLogStandardOutput:  Config.IsLogStandardOutput,
		LogFilePath:          Config.LogFilePath,
		IsExternalExtensions: Config.IsExternalExtensions,
	}
	if err := log.InitFromConfig("open-im-sdk-core", "",
		int(config.LogLevel), config.IsLogStandardOutput, false, config.LogFilePath,
		rotateCount, rotationTime); err != nil {
		f.respMessage.sendOnErrorResp(operationID, err)
		return
	}
	if f.userForSDK.InitSDK(config, callback) {
		f.respMessage.sendOnErrorResp(operationID, sdkerrs.ErrArgs)
	} else {
		f.respMessage.sendOnSuccessResp(operationID, "")
	}
}
func (f *FuncRouter) Login(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Login, args...)
}
