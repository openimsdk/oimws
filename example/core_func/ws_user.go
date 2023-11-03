package core_func

func (f *FuncRouter) GetUsersInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Full().GetUsersInfo, args...)
}

func (f *FuncRouter) GetUsersInfoWithCache(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Full().GetUsersInfoWithCache, args...)
}

func (f *FuncRouter) GetUsersInfoFromSrv(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().GetUsersInfo, args...)
}

func (f *FuncRouter) SetSelfInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().SetSelfInfo, args...)
}

func (f *FuncRouter) SetGlobalRecvMessageOpt(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().SetGlobalRecvMessageOpt, args...)
}

func (f *FuncRouter) GetSelfUserInfo(operationID string) {
	f.call(operationID, f.userForSDK.User().GetSelfUserInfo)
}

func (f *FuncRouter) UpdateMsgSenderInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().UpdateMsgSenderInfo, args...)
}

func (f *FuncRouter) SubscribeUsersStatus(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().SubscribeUsersStatus, args...)
}

func (f *FuncRouter) UnsubscribeUsersStatus(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().UnsubscribeUsersStatus, args...)
}

func (f *FuncRouter) GetSubscribeUsersStatus(operationID string) {
	f.call(operationID, f.userForSDK.User().GetSubscribeUsersStatus)
}

func (f *FuncRouter) GetUserStatus(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().GetUserStatus, args...)
}
