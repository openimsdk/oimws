package core_func

// UploadFile handles the file upload process.
func (f *FuncRouter) UploadFile(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.File().UploadFile, args)
}
