package core_func

// CreateGroup creates a new group with specified settings.
func (f *FuncRouter) CreateGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().CreateGroup, args...)
}

// JoinGroup sends a request to join a group.
func (f *FuncRouter) JoinGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().JoinGroup, args...)
}

// QuitGroup leaves a group.
func (f *FuncRouter) QuitGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().QuitGroup, args...)
}

// DismissGroup disbands a group.
func (f *FuncRouter) DismissGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().DismissGroup, args...)
}

// ChangeGroupMute toggles the mute status of the group.
func (f *FuncRouter) ChangeGroupMute(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().ChangeGroupMute, args...)
}

// ChangeGroupMemberMute toggles the mute status of a group member.
func (f *FuncRouter) ChangeGroupMemberMute(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().ChangeGroupMemberMute, args...)
}

// SetGroupMemberRoleLevel changes the role or level of a group member.
func (f *FuncRouter) SetGroupMemberRoleLevel(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupMemberRoleLevel, args...)
}

// SetGroupMemberInfo updates information for a group member.
func (f *FuncRouter) SetGroupMemberInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupMemberInfo, args...)
}

// GetJoinedGroupList retrieves the list of groups a user has joined.
func (f *FuncRouter) GetJoinedGroupList(operationID string) {
	f.call(operationID, f.userForSDK.Group().GetJoinedGroupList)
}

// GetSpecifiedGroupsInfo gets information of specified groups.
func (f *FuncRouter) GetSpecifiedGroupsInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().GetSpecifiedGroupsInfo, args...)
}

// SearchGroups searches for groups based on criteria.
func (f *FuncRouter) SearchGroups(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SearchGroups, args...)
}

// SetGroupInfo updates group settings or information.
func (f *FuncRouter) SetGroupInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupInfo, args...)
}

// SetGroupVerification sets the group's join request verification method.
func (f *FuncRouter) SetGroupVerification(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupVerification, args...)
}

// SetGroupLookMemberInfo toggles whether to allow group members to view each other's information.
func (f *FuncRouter) SetGroupLookMemberInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupLookMemberInfo, args...)
}

// SetGroupApplyMemberFriend configures settings related to adding group members as friends.
func (f *FuncRouter) SetGroupApplyMemberFriend(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupApplyMemberFriend, args...)
}

// GetGroupMemberList retrieves the member list of a group.
func (f *FuncRouter) GetGroupMemberList(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().GetGroupMemberList, args...)
}

// GetGroupMemberOwnerAndAdmin gets the owner and admin list of the group.
func (f *FuncRouter) GetGroupMemberOwnerAndAdmin(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().GetGroupMemberOwnerAndAdmin, args...)
}

// GetGroupMemberListByJoinTimeFilter gets the group member list filtered by join time.
func (f *FuncRouter) GetGroupMemberListByJoinTimeFilter(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().GetGroupMemberListByJoinTimeFilter, args...)
}

// GetSpecifiedGroupMembersInfo gets information for specified group members.
func (f *FuncRouter) GetSpecifiedGroupMembersInfo(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().GetSpecifiedGroupMembersInfo, args...)
}

// KickGroupMember removes a specific member from the group.
func (f *FuncRouter) KickGroupMember(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().KickGroupMember, args...)
}

// TransferGroupOwner changes the ownership of the group to another member.
func (f *FuncRouter) TransferGroupOwner(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().TransferGroupOwner, args...)
}

// InviteUserToGroup sends an invitation to a user to join the group.
func (f *FuncRouter) InviteUserToGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().InviteUserToGroup, args...)
}

// GetGroupApplicationListAsRecipient retrieves the list of join requests received by the group.
func (f *FuncRouter) GetGroupApplicationListAsRecipient(operationID string) {
	f.call(operationID, f.userForSDK.Group().GetGroupApplicationListAsRecipient)
}

// GetGroupApplicationListAsApplicant retrieves the list of join requests sent by the user.
func (f *FuncRouter) GetGroupApplicationListAsApplicant(operationID string) {
	f.call(operationID, f.userForSDK.Group().GetGroupApplicationListAsApplicant)
}

// AcceptGroupApplication approves a join request to the group.
func (f *FuncRouter) AcceptGroupApplication(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().AcceptGroupApplication, args...)
}

// RefuseGroupApplication denies a join request to the group.
func (f *FuncRouter) RefuseGroupApplication(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().RefuseGroupApplication, args...)
}

// SetGroupMemberNickname sets or changes a group member's nickname.
func (f *FuncRouter) SetGroupMemberNickname(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SetGroupMemberNickname, args...)
}

// SearchGroupMembers looks for members in the group matching certain criteria.
func (f *FuncRouter) SearchGroupMembers(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().SearchGroupMembers, args...)
}

// IsJoinGroup checks whether the user is a member of the specified group.
func (f *FuncRouter) IsJoinGroup(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Group().IsJoinGroup, args...)
}
