package core_func

func (f *FuncRouter) GetAllConversationList(operationID string) {
	f.call(operationID, f.userForSDK.Conversation().GetAllConversationList)
}

func (f *FuncRouter) GetConversationListSplit(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetConversationListSplit, args...)
}

func (f *FuncRouter) GetOneConversation(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetOneConversation, args...)
}

func (f *FuncRouter) GetMultipleConversation(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetMultipleConversation, args...)
}

func (f *FuncRouter) SetConversationMsgDestructTime(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetConversationMsgDestructTime, args...)
}

func (f *FuncRouter) SetConversationIsMsgDestruct(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetConversationIsMsgDestruct, args...)
}

func (f *FuncRouter) HideConversation(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().HideConversation, args...)
}

// deprecated
func (f *FuncRouter) GetConversationRecvMessageOpt(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetConversationRecvMessageOpt, args...)
}

func (f *FuncRouter) SetConversationDraft(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetConversationDraft, args...)
}

func (f *FuncRouter) ResetConversationGroupAtType(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().ResetConversationGroupAtType, args...)
}

func (f *FuncRouter) PinConversation(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().PinConversation, args...)
}

func (f *FuncRouter) SetConversationPrivateChat(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetOneConversationPrivateChat, args...)
}

func (f *FuncRouter) SetConversationBurnDuration(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetOneConversationBurnDuration, args...)
}

func (f *FuncRouter) SetConversationRecvMessageOpt(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetOneConversationRecvMessageOpt, args...)
}

func (f *FuncRouter) GetTotalUnreadMsgCount(operationID string) {
	f.call(operationID, f.userForSDK.Conversation().GetTotalUnreadMsgCount)
}

func (f *FuncRouter) SendMessage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SendMessage, args)
}
func (f *FuncRouter) SendMessageNotOss(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SendMessageNotOss, args)
}

// deprecated
func (f *FuncRouter) SendMessageByBuffer(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SendMessageByBuffer, args)
}

func (f *FuncRouter) FindMessageList(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().FindMessageList, args...)
}

func (f *FuncRouter) GetAdvancedHistoryMessageList(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetAdvancedHistoryMessageList, args...)
}

func (f *FuncRouter) GetAdvancedHistoryMessageListReverse(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetAdvancedHistoryMessageListReverse, args...)
}

func (f *FuncRouter) RevokeMessage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().RevokeMessage, args...)
}

func (f *FuncRouter) TypingStatusUpdate(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().TypingStatusUpdate, args...)
}

func (f *FuncRouter) MarkConversationMessageAsRead(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().MarkConversationMessageAsRead, args...)
}

func (f *FuncRouter) MarkMessagesAsReadByMsgID(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().MarkMessagesAsReadByMsgID, args...)
}

func (f *FuncRouter) DeleteMessageFromLocalStorage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().DeleteMessageFromLocalStorage, args...)
}

func (f *FuncRouter) DeleteMessage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().DeleteMessage, args...)
}

func (f *FuncRouter) HideAllConversations(operationID string) {
	f.call(operationID, f.userForSDK.Conversation().HideAllConversations)
}

func (f *FuncRouter) DeleteAllMsgFromLocalAndSvr(operationID string) {
	f.call(operationID, f.userForSDK.Conversation().DeleteAllMsgFromLocalAndSvr)
}

func (f *FuncRouter) DeleteAllMsgFromLocal(operationID string) {
	f.call(operationID, f.userForSDK.Conversation().DeleteAllMessageFromLocalStorage)
}

func (f *FuncRouter) ClearConversationAndDeleteAllMsg(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().ClearConversationAndDeleteAllMsg, args...)
}

func (f *FuncRouter) DeleteConversationAndDeleteAllMsg(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().DeleteConversationAndDeleteAllMsg, args...)
}

func (f *FuncRouter) InsertSingleMessageToLocalStorage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().InsertSingleMessageToLocalStorage, args...)
}

func (f *FuncRouter) InsertGroupMessageToLocalStorage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().InsertGroupMessageToLocalStorage, args...)
}

func (f *FuncRouter) SearchLocalMessages(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SearchLocalMessages, args...)
}

func (f *FuncRouter) SetMessageLocalEx(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetMessageLocalEx, args...)
}
