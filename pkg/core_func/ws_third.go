package core_func

// UpdateFcmToken updates the FCM (Firebase Cloud Messaging) token for the current device session.
func (f *FuncRouter) UpdateFcmToken(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Third().UpdateFcmToken, args...)
}

// SetAppBadge sets the badge count for the application's icon, typically reflecting the number of unread notifications or messages.
func (f *FuncRouter) SetAppBadge(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Third().SetAppBadge, args...)
}

// UploadLogs initiates the upload of application logs to a remote server for diagnostic purposes.
func (f *FuncRouter) UploadLogs(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Third().UploadLogs, args...)
}
