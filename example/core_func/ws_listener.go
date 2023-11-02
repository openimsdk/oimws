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

// 获取自身函数名称
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
}
