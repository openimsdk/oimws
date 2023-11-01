package core_func

import (
	"encoding/json"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/sdkerrs"
	"github.com/openimsdk/openim-sdk-core/v3/sdk_struct"
)

func (f *FuncRouter) InitSDK(operationID string, args ...any) {
	callback := NewConnCallback(f.respMessage)
	if len(args) == 0 {
		f.respMessage.sendOnErrorResp(operationID, sdkerrs.ErrArgs)
	}
	config := sdk_struct.IMConfig{}
	if v, ok := args[0].(string); ok {
		if err := json.Unmarshal([]byte(v), &config); err != nil {
			f.respMessage.sendOnErrorResp(operationID, sdkerrs.ErrArgs)
		}
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
