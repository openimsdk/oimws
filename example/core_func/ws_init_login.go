package core_func

import (
	"fmt"
	"strconv"

	"github.com/OpenIMSDK/tools/log"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/sdkerrs"
	"github.com/openimsdk/openim-sdk-core/v3/sdk_struct"
)

var Config sdk_struct.IMConfig

const (
	rotateCount  uint = 0
	rotationTime uint = 24
)

// InitSDK initializes the SDK with the given operation ID and platform ID.
func (f *FuncRouter) InitSDK(operationID, platformID string) {
	fmt.Println("InitSDK", "data=", platformID, operationID)
	callback := NewConnCallback(f.respMessage)
	j, err := strconv.ParseInt(platformID, 10, 64)
	if err != nil {
		f.respMessage.sendOnErrorResp(operationID, "InitSDK", err)
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
		f.respMessage.sendOnErrorResp(operationID, "InitSDK", err)
		return
	}
	if f.userForSDK.InitSDK(config, callback) {
		f.respMessage.sendOnSuccessResp(operationID, "InitSDK", "")
	} else {
		f.respMessage.sendOnErrorResp(operationID, "InitSDK", sdkerrs.ErrArgs)
	}
}

// UnInitSDK uninitializes the SDK.
func (f *FuncRouter) UnInitSDK(operationID string) {
	if f.userForSDK == nil {
		fmt.Println(operationID, "UserForSDK is nil,")
		return
	}
	f.userForSDK.UnInitSDK()
	f.userForSDK = nil

}

// Login logs in a user using the provided arguments.
func (f *FuncRouter) Login(operationID string, args ...any) {
	f.setAllListener()
	f.call(operationID, f.userForSDK.Login, args...)
}

// Logout logs out the current user.
func (f *FuncRouter) Logout(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Logout, args...)
}

// GetLoginUserID returns the logged-in user's ID.
func (f *FuncRouter) GetLoginUserID() string {
	if f.userForSDK == nil {
		return ""
	}
	return f.userForSDK.GetLoginUserID()
}

// SetAppBackgroundStatus updates the app's background status.
func (f *FuncRouter) SetAppBackgroundStatus(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.SetAppBackgroundStatus, args...)
}

// NetworkStatusChanged handles the change in network status.
func (f *FuncRouter) NetworkStatusChanged(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.NetworkStatusChanged, args...)
}

// GetLoginStatus retrieves the current login status of the user.
func (f *FuncRouter) GetLoginStatus(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.GetLoginStatus, args...)
}

// setAllListener sets all listeners for the SDK to handle various events.
func (f *FuncRouter) setAllListener() {
	f.userForSDK.SetConversationListener(NewConversationCallback(f.respMessage))
	f.userForSDK.SetGroupListener(NewGroupCallback(f.respMessage))
	f.userForSDK.SetUserListener(NewUserCallback(f.respMessage))
	f.userForSDK.SetAdvancedMsgListener(NewAdvancedMsgCallback(f.respMessage))
	f.userForSDK.SetFriendListener(NewFriendCallback(f.respMessage))
	f.userForSDK.SetBatchMsgListener(NewBatchMessageCallback(f.respMessage))

}
