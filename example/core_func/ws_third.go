package core_func

func (f *FuncRouter) UpdateFcmToken(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Third().UpdateFcmToken, args...)
}

func (f *FuncRouter) SetAppBadge(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Third().SetAppBadge, args...)
}

func (f *FuncRouter) UploadLogs(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Third().UploadLogs, args...)
}
