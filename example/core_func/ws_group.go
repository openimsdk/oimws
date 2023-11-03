package core_func

func (f *FuncRouter) CreateGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().CreateGroup, args...)
}

func (f *FuncRouter) JoinGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().JoinGroup, args...)
}

func (f *FuncRouter) QuitGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().QuitGroup, args...)
}

func (f *FuncRouter) DismissGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().DismissGroup, args...)
}

func (f *FuncRouter) ChangeGroupMute(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().ChangeGroupMute, args...)
}

func (f *FuncRouter) ChangeGroupMemberMute(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().ChangeGroupMemberMute, args...)
}

func (f *FuncRouter) SetGroupMemberRoleLevel(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupMemberRoleLevel, args...)
}

func (f *FuncRouter) SetGroupMemberInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupMemberInfo, args...)
}

func (f *FuncRouter) GetJoinedGroupList(operationID string) {
	f.call(operationID, f.userForSDK.Group().GetJoinedGroupList)
}

func (f *FuncRouter) GetSpecifiedGroupsInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().GetSpecifiedGroupsInfo, args...)
}

func (f *FuncRouter) SearchGroups(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SearchGroups, args...)
}

func (f *FuncRouter) SetGroupInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupInfo, args...)
}

func (f *FuncRouter) SetGroupVerification(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupVerification, args...)
}

func (f *FuncRouter) SetGroupLookMemberInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupLookMemberInfo, args...)
}

func (f *FuncRouter) SetGroupApplyMemberFriend(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupApplyMemberFriend, args...)
}

func (f *FuncRouter) GetGroupMemberList(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().GetGroupMemberList, args...)
}

func (f *FuncRouter) GetGroupMemberOwnerAndAdmin(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().GetGroupMemberOwnerAndAdmin, args...)
}

func (f *FuncRouter) GetGroupMemberListByJoinTimeFilter(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().GetGroupMemberListByJoinTimeFilter, args...)
}

func (f *FuncRouter) GetSpecifiedGroupMembersInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().GetSpecifiedGroupMembersInfo, args...)
}

func (f *FuncRouter) KickGroupMember(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().KickGroupMember, args...)
}

func (f *FuncRouter) TransferGroupOwner(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().TransferGroupOwner, args...)
}

func (f *FuncRouter) InviteUserToGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().InviteUserToGroup, args...)
}

func (f *FuncRouter) GetGroupApplicationListAsRecipient(operationID string) {
	f.call(operationID, f.userForSDK.Group().GetGroupApplicationListAsRecipient)
}

func (f *FuncRouter) GetGroupApplicationListAsApplicant(operationID string) {
	f.call(operationID, f.userForSDK.Group().GetGroupApplicationListAsApplicant)
}

func (f *FuncRouter) AcceptGroupApplication(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().AcceptGroupApplication, args...)
}

func (f *FuncRouter) RefuseGroupApplication(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().RefuseGroupApplication, args...)
}

func (f *FuncRouter) SetGroupMemberNickname(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupMemberNickname, args...)
}

func (f *FuncRouter) SearchGroupMembers(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SearchGroupMembers, args...)
}

func (f *FuncRouter) IsJoinGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().IsJoinGroup, args...)
}
