package core_func

func (f *FuncRouter) CheckFriend(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().CheckFriend, args...)
}

func (f *FuncRouter) GetSpecifiedFriendsInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().GetSpecifiedFriendsInfo, args...)
}

func (f *FuncRouter) GetFriendList(operationID string) {
	f.call(operationID, f.userForSDK.Friend().GetFriendList)
}

func (f *FuncRouter) GetFriendListPage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().GetFriendListPage, args...)
}

func (f *FuncRouter) SearchFriends(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().SearchFriends, args...)
}

func (f *FuncRouter) AddFriend(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().AddFriend, args...)
}

func (f *FuncRouter) SetFriendRemark(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().SetFriendRemark, args...)
}

func (f *FuncRouter) DeleteFriend(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().DeleteFriend, args...)
}

func (f *FuncRouter) GetFriendApplicationListAsRecipient(operationID string) {
	f.call(operationID, f.userForSDK.Friend().GetFriendApplicationListAsRecipient)
}

func (f *FuncRouter) GetFriendApplicationListAsApplicant(operationID string) {
	f.call(operationID, f.userForSDK.Friend().GetFriendApplicationListAsApplicant)
}

func (f *FuncRouter) AcceptFriendApplication(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().AcceptFriendApplication, args...)
}

func (f *FuncRouter) RefuseFriendApplication(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().RefuseFriendApplication, args...)
}

func (f *FuncRouter) AddBlack(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().AddBlack, args...)
}

func (f *FuncRouter) GetBlackList(operationID string) {
	f.call(operationID, f.userForSDK.Friend().GetBlackList)
}

func (f *FuncRouter) RemoveBlack(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().RemoveBlack, args...)
}
