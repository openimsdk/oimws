package core_func

import (
	"fmt"
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
	fmt.Println("InitSDK", "data=", platformID, operationID)
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
		f.respMessage.sendOnSuccessResp(operationID, "")
	} else {
		f.respMessage.sendOnErrorResp(operationID, sdkerrs.ErrArgs)
	}
}

func (f *FuncRouter) UnInitSDK(operationID string) {
	if f.userForSDK == nil {
		fmt.Println(operationID, "UserForSDK is nil,")
		return
	}
	f.userForSDK.UnInitSDK()
	f.userForSDK = nil

}

func (f *FuncRouter) Login(operationID string, args ...any) {
	f.setAllListener()
	fmt.Println(operationID, "Login")
	f.call(operationID, f.userForSDK.Login, args...)
}

func (f *FuncRouter) Logout(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Logout, args...)
}

func (f *FuncRouter) GetLoginUserID() string {
	if f.userForSDK == nil {
		return ""
	}
	return f.userForSDK.GetLoginUserID()
}

func (f *FuncRouter) SetAppBackgroundStatus(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.SetAppBackgroundStatus, args...)
}
func (f *FuncRouter) NetworkStatusChanged(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.NetworkStatusChanged, args...)
}
func (f *FuncRouter) GetLoginStatus(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.GetLoginStatus, args...)
}
func (f *FuncRouter) setAllListener() {
	f.userForSDK.SetConversationListener(NewConversationCallback(f.respMessage))
	f.userForSDK.SetGroupListener(NewGroupCallback(f.respMessage))
	f.userForSDK.SetUserListener(NewUserCallback(f.respMessage))
	f.userForSDK.SetAdvancedMsgListener(NewAdvancedMsgCallback(f.respMessage))
	f.userForSDK.SetFriendListener(NewFriendCallback(f.respMessage))

}
