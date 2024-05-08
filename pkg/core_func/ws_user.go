package core_func

// GetUsersInfo retrieves information for the specified users.
func (f *FuncRouter) GetUsersInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Full().GetUsersInfo, args...)
}

// GetUsersInfoWithCache retrieves information for the specified users and uses local cache if available.
func (f *FuncRouter) GetUsersInfoWithCache(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Full().GetUsersInfoWithCache, args...)
}

// GetUsersInfoFromSrv retrieves user information directly from the server, bypassing any cache.
func (f *FuncRouter) GetUsersInfoFromSrv(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().GetUsersInfo, args...)
}

// SetSelfInfo updates the current user's information.
func (f *FuncRouter) SetSelfInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().SetSelfInfo, args...)
}

// SetSelfInfoEx updates the current user's information.
func (f *FuncRouter) SetSelfInfoEx(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().SetSelfInfo, args...)
}

// SetGlobalRecvMessageOpt sets the global option for receiving messages for the user.
func (f *FuncRouter) SetGlobalRecvMessageOpt(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().SetGlobalRecvMessageOpt, args...)
}

// GetSelfUserInfo retrieves the current user's information.
func (f *FuncRouter) GetSelfUserInfo(operationID string) {
	f.call(operationID, f.userForSDK.User().GetSelfUserInfo)
}

// UpdateMsgSenderInfo updates the information of the message sender.
func (f *FuncRouter) UpdateMsgSenderInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().UpdateMsgSenderInfo, args...)
}

// SubscribeUsersStatus subscribes to the online status updates of the specified users.
func (f *FuncRouter) SubscribeUsersStatus(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().SubscribeUsersStatus, args...)
}

// UnsubscribeUsersStatus unsubscribes from the online status updates of the specified users.
func (f *FuncRouter) UnsubscribeUsersStatus(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().UnsubscribeUsersStatus, args...)
}

// GetSubscribeUsersStatus retrieves the subscription status of online status updates for the specified users.
func (f *FuncRouter) GetSubscribeUsersStatus(operationID string) {
	f.call(operationID, f.userForSDK.User().GetSubscribeUsersStatus)
}

// GetUserStatus retrieves the current status of a user.
func (f *FuncRouter) GetUserStatus(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.User().GetUserStatus, args...)
}
