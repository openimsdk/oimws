package core_func

func (f *FuncRouter) CheckFriend(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().CheckFriend, args...)
}
