package core_func

// CheckFriend checks the friend status between users.
func (f *FuncRouter) CheckFriend(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().CheckFriend, args...)
}

// GetSpecifiedFriendsInfo gets the information of specific friends.
func (f *FuncRouter) GetSpecifiedFriendsInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().GetSpecifiedFriendsInfo, args...)
}

// GetFriendList retrieves the user's friend list.
func (f *FuncRouter) GetFriendList(operationID string) {
	f.call(operationID, f.userForSDK.Friend().GetFriendList)
}

// GetFriendListPage retrieves a paginated friend list.
func (f *FuncRouter) GetFriendListPage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().GetFriendListPage, args...)
}

// SearchFriends searches for friends based on given criteria.
func (f *FuncRouter) SearchFriends(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().SearchFriends, args...)
}

// AddFriend sends a friend request to another user.
func (f *FuncRouter) AddFriend(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().AddFriend, args...)
}

// SetFriendRemark sets a remark for a friend.
func (f *FuncRouter) SetFriendRemark(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().SetFriendRemark, args...)
}

// PinFriends pins friends to the top of the friend list.
func (f *FuncRouter) PinFriends(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().PinFriends, args...)
}

// DeleteFriend removes a user from the friend list.
func (f *FuncRouter) DeleteFriend(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().DeleteFriend, args...)
}

// GetFriendApplicationListAsRecipient retrieves friend requests received.
func (f *FuncRouter) GetFriendApplicationListAsRecipient(operationID string) {
	f.call(operationID, f.userForSDK.Friend().GetFriendApplicationListAsRecipient)
}

// GetFriendApplicationListAsApplicant retrieves friend requests sent.
func (f *FuncRouter) GetFriendApplicationListAsApplicant(operationID string) {
	f.call(operationID, f.userForSDK.Friend().GetFriendApplicationListAsApplicant)
}

// AcceptFriendApplication accepts a friend request.
func (f *FuncRouter) AcceptFriendApplication(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().AcceptFriendApplication, args...)
}

// RefuseFriendApplication declines a friend request.
func (f *FuncRouter) RefuseFriendApplication(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().RefuseFriendApplication, args...)
}

// AddBlack adds a user to the blacklist.
func (f *FuncRouter) AddBlack(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().AddBlack, args...)
}

// GetBlackList retrieves the blacklist.
func (f *FuncRouter) GetBlackList(operationID string) {
	f.call(operationID, f.userForSDK.Friend().GetBlackList)
}

// RemoveBlack removes a user from the blacklist.
func (f *FuncRouter) RemoveBlack(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().RemoveBlack, args...)
}

// SetFriendsEx sets the friend ex info.
func (f *FuncRouter) SetFriendsEx(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Friend().SetFriendsEx, args...)
}
