package core_func

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

/*
	连接状态相关的内容
*/

type ConnCallback struct {
	respMessage *RespMessage
}

func NewConnCallback(respMessage *RespMessage) *ConnCallback {
	return &ConnCallback{respMessage: respMessage}
}

// OnConnecting 连接中
func (c ConnCallback) OnConnecting() {
	c.respMessage.sendEventSuccessRespNoData(getSelfFuncName())
}

// OnConnectSuccess 连接成功
func (c ConnCallback) OnConnectSuccess() {
	c.respMessage.sendEventSuccessRespNoData(getSelfFuncName())
}

// OnConnectFailed 连接失败
func (c ConnCallback) OnConnectFailed(errCode int32, errMsg string) {
	c.respMessage.sendEventFailedRespNoData(getSelfFuncName(), errCode, errMsg)
}

// OnKickedOffline 强制下线
func (c ConnCallback) OnKickedOffline() {
	c.respMessage.sendEventSuccessRespNoData(getSelfFuncName())
}

// OnUserTokenExpired token 过期
func (c ConnCallback) OnUserTokenExpired() {
	c.respMessage.sendEventSuccessRespNoData(getSelfFuncName())
}

// getSelfFuncName gets the name of the caller function.
// This would need to be implemented or replaced with appropriate functionality.
func getSelfFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	end := strings.LastIndex(runtime.FuncForPC(pc).Name(), ".")
	if end == -1 {
		return ""
	}
	return runtime.FuncForPC(pc).Name()[end+1:]
}

type ConversationCallback struct {
	respMessage *RespMessage
}

func NewConversationCallback(respMessage *RespMessage) *ConversationCallback {
	return &ConversationCallback{respMessage: respMessage}
}

func (c ConversationCallback) OnSyncServerStart() {
	c.respMessage.sendEventSuccessRespNoData(getSelfFuncName())
}

func (c ConversationCallback) OnSyncServerFinish() {
	c.respMessage.sendEventSuccessRespNoData(getSelfFuncName())
}

func (c ConversationCallback) OnSyncServerFailed() {
	c.respMessage.sendEventFailedREspNoErr(getSelfFuncName())
}

func (c ConversationCallback) OnNewConversation(conversationList string) {
	c.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), conversationList)
}

func (c ConversationCallback) OnConversationChanged(conversationList string) {
	c.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), conversationList)
}

func (c ConversationCallback) OnTotalUnreadMessageCountChanged(totalUnreadCount int32) {
	c.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), fmt.Sprintf("%d", totalUnreadCount))
}

type AdvancedMsgCallback struct {
	respMessage *RespMessage
}

func NewAdvancedMsgCallback(respMessage *RespMessage) *AdvancedMsgCallback {
	return &AdvancedMsgCallback{respMessage: respMessage}
}

func (a AdvancedMsgCallback) OnRecvNewMessage(message string) {
	a.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), message)
}

func (a AdvancedMsgCallback) OnRecvC2CReadReceipt(msgReceiptList string) {
	a.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), msgReceiptList)
}

func (a AdvancedMsgCallback) OnRecvGroupReadReceipt(groupMsgReceiptList string) {
	a.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupMsgReceiptList)
}

func (a AdvancedMsgCallback) OnRecvMessageRevoked(msgID string) {
	a.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), msgID)
}

func (a AdvancedMsgCallback) OnNewRecvMessageRevoked(messageRevoked string) {
	a.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), messageRevoked)
}

func (a AdvancedMsgCallback) OnRecvMessageModified(message string) {
	a.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), message)
}

func (a AdvancedMsgCallback) OnRecvMessageExtensionsChanged(clientMsgID string, reactionExtensionList string) {
	m := make(map[string]interface{})
	m["clientMsgID"] = clientMsgID
	m["reactionExtensionList"] = reactionExtensionList
	dataType, _ := json.Marshal(m)
	dataString := string(dataType)
	a.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), dataString)
}

func (a AdvancedMsgCallback) OnRecvMessageExtensionsDeleted(clientMsgID string, reactionExtensionKeyList string) {
	m := make(map[string]interface{})
	m["clientMsgID"] = clientMsgID
	m["reactionExtensionKeyList"] = reactionExtensionKeyList
	dataType, _ := json.Marshal(m)
	dataString := string(dataType)
	a.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), dataString)
}

func (a AdvancedMsgCallback) OnRecvMessageExtensionsAdded(clientMsgID string, reactionExtensionList string) {
	m := make(map[string]interface{})
	m["clientMsgID"] = clientMsgID
	m["reactionExtensionList"] = reactionExtensionList
	dataType, _ := json.Marshal(m)
	dataString := string(dataType)
	a.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), dataString)
}

func (a AdvancedMsgCallback) OnRecvOfflineNewMessage(message string) {
	a.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), message)
}

func (a AdvancedMsgCallback) OnMsgDeleted(message string) {
	a.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), message)
}

type BaseCallback struct {
	respMessage *RespMessage
}d

type BatchMessageCallback struct {
	respMessage *RespMessage
}

// NewBatchMessageCallback creates a new instance of BatchMessageCallback.
func NewBatchMessageCallback(respMessage *RespMessage) *BatchMessageCallback {
	return &BatchMessageCallback{respMessage: respMessage}
}

// OnRecvNewMessages is called when new messages are received.
func (b *BatchMessageCallback) OnRecvNewMessages(messageList string) {
	b.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), messageList)
}

// OnRecvOfflineNewMessages is called when new offline messages are received.
func (b *BatchMessageCallback) OnRecvOfflineNewMessages(messageList string) {
	b.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), messageList)
}


type FriendCallback struct {
	respMessage *RespMessage
}

// NewFriendCallback creates a new instance of FriendCallback.
func NewFriendCallback(respMessage *RespMessage) *FriendCallback {
	return &FriendCallback{respMessage: respMessage}
}

// OnFriendApplicationAdded notifies when a friend application is added.
func (f *FriendCallback) OnFriendApplicationAdded(friendApplication string) {
	f.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), friendApplication)
}

// OnFriendApplicationDeleted notifies when a friend application is deleted.
func (f *FriendCallback) OnFriendApplicationDeleted(friendApplication string) {
	f.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), friendApplication)
}

// OnFriendApplicationAccepted notifies when a friend application is accepted.
func (f *FriendCallback) OnFriendApplicationAccepted(friendApplication string) {
	f.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), friendApplication)
}

// OnFriendApplicationRejected notifies when a friend application is rejected.
func (f *FriendCallback) OnFriendApplicationRejected(friendApplication string) {
	f.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), friendApplication)
}

// OnFriendAdded notifies when a new friend is added.
func (f *FriendCallback) OnFriendAdded(friendInfo string) {
	f.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), friendInfo)
}

// OnFriendDeleted notifies when a friend is deleted.
func (f *FriendCallback) OnFriendDeleted(friendInfo string) {
	f.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), friendInfo)
}

// OnFriendInfoChanged notifies when friend information is changed.
func (f *FriendCallback) OnFriendInfoChanged(friendInfo string) {
	f.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), friendInfo)
}

// OnBlackAdded notifies when a black list entry is added.
func (f *FriendCallback) OnBlackAdded(blackInfo string) {
	f.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), blackInfo)
}

// OnBlackDeleted notifies when a black list entry is deleted.
func (f *FriendCallback) OnBlackDeleted(blackInfo string) {
	f.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), blackInfo)
}


type GroupCallback struct {
	respMessage *RespMessage
}

// NewGroupCallback creates a new instance of GroupCallback.
func NewGroupCallback(respMessage *RespMessage) *GroupCallback {
	return &GroupCallback{respMessage: respMessage}
}

// OnJoinedGroupAdded notifies the client that a group has been joined.
func (g *GroupCallback) OnJoinedGroupAdded(groupInfo string) {
	g.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupInfo)
}

// OnJoinedGroupDeleted notifies the client that a joined group has been deleted.
func (g *GroupCallback) OnJoinedGroupDeleted(groupInfo string) {
	g.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupInfo)
}

// OnGroupMemberAdded notifies the client that a new member has been added to a group.
func (g *GroupCallback) OnGroupMemberAdded(groupMemberInfo string) {
	g.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupMemberInfo)
}

// OnGroupMemberDeleted notifies the client that a member has been removed from a group.
func (g *GroupCallback) OnGroupMemberDeleted(groupMemberInfo string) {
	g.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupMemberInfo)
}

// OnGroupApplicationAdded notifies the client that a group application has been received.
func (g *GroupCallback) OnGroupApplicationAdded(groupApplication string) {
	g.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupApplication)
}

// OnGroupApplicationDeleted notifies the client that a group application has been deleted.
func (g *GroupCallback) OnGroupApplicationDeleted(groupApplication string) {
	g.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupApplication)
}

// OnGroupInfoChanged notifies the client that group information has changed.
func (g *GroupCallback) OnGroupInfoChanged(groupInfo string) {
	g.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupInfo)
}

// OnGroupMemberInfoChanged notifies the client that group member information has changed.
func (g *GroupCallback) OnGroupMemberInfoChanged(groupMemberInfo string) {
	g.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupMemberInfo)
}

// OnGroupApplicationAccepted notifies the client that a group application has been accepted.
func (g *GroupCallback) OnGroupApplicationAccepted(groupApplication string) {
	g.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupApplication)
}

// OnGroupApplicationRejected notifies the client that a group application has been rejected.
func (g *GroupCallback) OnGroupApplicationRejected(groupApplication string) {
	g.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupApplication)
}

// OnGroupDismissed notifies the client that a group has been dismissed.
func (g *GroupCallback) OnGroupDismissed(groupInfo string) {
	g.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), groupInfo)
}


// UserCallback represents a callback handler for user-related events.
type UserCallback struct {
	respMessage *RespMessage
}

// NewUserCallback creates a new UserCallback handler.
func NewUserCallback(respMessage *RespMessage) *UserCallback {
	return &UserCallback{respMessage: respMessage}
}

// OnUserStatusChanged is triggered when there is a change in the user status.
func (u *UserCallback) OnUserStatusChanged(statusMap string) {
	u.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), statusMap)
}

// OnSelfInfoUpdated is triggered when the user's own information is updated.
func (u *UserCallback) OnSelfInfoUpdated(userInfo string) {
	u.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), userInfo)
}

type CustomBusinessCallback struct {
	respMessage *RespMessage
}

// NewCustomBusinessCallback creates a new instance of CustomBusinessCallback with the provided RespMessage.
func NewCustomBusinessCallback(respMessage *RespMessage) *CustomBusinessCallback {
	return &CustomBusinessCallback{respMessage: respMessage}
}

// OnRecvCustomBusinessMessage is called when a custom business message is received.
func (cb *CustomBusinessCallback) OnRecvCustomBusinessMessage(businessMessage string) {
	cb.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), businessMessage)
}

type SignalingCallback struct {
	respMessage *RespMessage
}

// NewSignalingCallback creates a new instance of SignalingCallback with the provided RespMessage.
func NewSignalingCallback(respMessage *RespMessage) *SignalingCallback {
	return &SignalingCallback{respMessage: respMessage}
}

// OnRoomParticipantConnected is called when a room participant successfully connects.
func (sc *SignalingCallback) OnRoomParticipantConnected(participantConnectedData string) {
	sc.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), participantConnectedData)
}

// OnRoomParticipantDisconnected is called when a room participant gets disconnected.
func (sc *SignalingCallback) OnRoomParticipantDisconnected(participantDisconnectedData string) {
	sc.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), participantDisconnectedData)
}

// OnReceiveNewInvitation is called when a new invitation is received.
func (sc *SignalingCallback) OnReceiveNewInvitation(newInvitationData string) {
	sc.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), newInvitationData)
}

// OnInviteeAccepted is called when an invitee accepts an invitation.
func (sc *SignalingCallback) OnInviteeAccepted(acceptedData string) {
	sc.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), acceptedData)
}

// OnInviteeAcceptedByOtherDevice is called when an invitee accepts an invitation from another device.
func (sc *SignalingCallback) OnInviteeAcceptedByOtherDevice(acceptedData string) {
	sc.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), acceptedData)
}

// OnInviteeRejected is called when an invitee rejects an invitation.
func (sc *SignalingCallback) OnInviteeRejected(rejectedData string) {
	sc.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), rejectedData)
}

// OnInviteeRejectedByOtherDevice is called when an invitee rejects an invitation from another device.
func (sc *SignalingCallback) OnInviteeRejectedByOtherDevice(rejectedData string) {
	sc.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), rejectedData)
}

// OnInvitationCancelled is called when an invitation is cancelled.
func (sc *SignalingCallback) OnInvitationCancelled(cancelledData string) {
	sc.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), cancelledData)
}

// OnInvitationTimeout is called when an invitation times out.
func (sc *SignalingCallback) OnInvitationTimeout(timeoutData string) {
	sc.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), timeoutData)
}

// OnHangUp is called when a hang-up event occurs.
func (sc *SignalingCallback) OnHangUp(hangUpData string) {
	sc.respMessage.sendEventSuccessRespWithData(getSelfFuncName(), hangUpData)
}
