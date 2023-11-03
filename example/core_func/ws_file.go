package core_func

func (f *FuncRouter) UploadFile(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.File().UploadFile, args)
}
