package core_func

// GetAllConversationList retrieves all conversations associated with a user.
func (f *FuncRouter) GetAllConversationList(operationID string) {
	f.call(operationID, f.userForSDK.Conversation().GetAllConversationList)
}

// GetConversationListSplit retrieves a portion of the user's conversation list, based on provided criteria.
func (f *FuncRouter) GetConversationListSplit(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetConversationListSplit, args...)
}

// GetOneConversation fetches a single conversation based on specified identifiers such as conversation ID.
func (f *FuncRouter) GetOneConversation(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetOneConversation, args...)
}

// GetMultipleConversation retrieves multiple conversations, usually filtered by a set of identifiers.
func (f *FuncRouter) GetMultipleConversation(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetMultipleConversation, args...)
}

// SetConversationMsgDestructTime sets a timer after which messages in the conversation will be destroyed.
func (f *FuncRouter) SetConversationMsgDestructTime(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetConversationMsgDestructTime, args...)
}

// SetConversationIsMsgDestruct toggles the self-destruction feature for messages in a conversation.
func (f *FuncRouter) SetConversationIsMsgDestruct(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetConversationIsMsgDestruct, args...)
}

// HideConversation hides a conversation from the conversation list without deleting it.
func (f *FuncRouter) HideConversation(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().HideConversation, args...)
}

// GetConversationRecvMessageOpt retrieves the options for receiving messages in a conversation.
// deprecated.
func (f *FuncRouter) GetConversationRecvMessageOpt(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetConversationRecvMessageOpt, args...)
}

// SetConversationDraft saves a draft message in a conversation.
func (f *FuncRouter) SetConversationDraft(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetConversationDraft, args...)
}

// ResetConversationGroupAtType resets the notification state for when a user is mentioned in a group conversation.
func (f *FuncRouter) ResetConversationGroupAtType(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().ResetConversationGroupAtType, args...)
}

// PinConversation pins a conversation to the top of the conversation list for quick access.
func (f *FuncRouter) PinConversation(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().PinConversation, args...)
}

// SetConversationPrivateChat sets a conversation as a private chat, likely with enhanced privacy settings.
func (f *FuncRouter) SetConversationPrivateChat(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetOneConversationPrivateChat, args...)
}

// SetConversationBurnDuration sets the duration before messages in a private chat are automatically deleted.
func (f *FuncRouter) SetConversationBurnDuration(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetOneConversationBurnDuration, args...)
}

// SetConversationRecvMessageOpt sets options for how messages should be received in a conversation.
func (f *FuncRouter) SetConversationRecvMessageOpt(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetOneConversationRecvMessageOpt, args...)
}

// GetTotalUnreadMsgCount gets the count of all unread messages across all conversations for a user.
func (f *FuncRouter) GetTotalUnreadMsgCount(operationID string) {
	f.call(operationID, f.userForSDK.Conversation().GetTotalUnreadMsgCount)
}

// SendMessage sends a message within a conversation.
func (f *FuncRouter) SendMessage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SendMessage, args)
}

// SendMessageNotOss sends a message without using an object storage service for any attachments.
func (f *FuncRouter) SendMessageNotOss(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SendMessageNotOss, args)
}

// SendMessageByBuffer sends a message using a buffer.
func (f *FuncRouter) SendMessageByBuffer(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SendMessageByBuffer, args)
}

// FindMessageList retrieves a list of messages based on search criteria.
func (f *FuncRouter) FindMessageList(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().FindMessageList, args...)
}

// GetAdvancedHistoryMessageList retrieves a historical list of messages with advanced filtering options.
func (f *FuncRouter) GetAdvancedHistoryMessageList(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetAdvancedHistoryMessageList, args...)
}

// GetAdvancedHistoryMessageListReverse retrieves a historical list of messages in reverse order with advanced filtering.
func (f *FuncRouter) GetAdvancedHistoryMessageListReverse(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().GetAdvancedHistoryMessageListReverse, args...)
}

// RevokeMessage revokes or recalls a message that was previously sent.
func (f *FuncRouter) RevokeMessage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().RevokeMessage, args...)
}

// TypingStatusUpdate sends an indication that the user is typing a message.
func (f *FuncRouter) TypingStatusUpdate(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().TypingStatusUpdate, args...)
}

// MarkConversationMessageAsRead marks all messages in a conversation as read.
func (f *FuncRouter) MarkConversationMessageAsRead(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().MarkConversationMessageAsRead, args...)
}

// MarkMessagesAsReadByMsgID marks specific messages as read using their message IDs.
func (f *FuncRouter) MarkMessagesAsReadByMsgID(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().MarkMessagesAsReadByMsgID, args...)
}

// DeleteMessageFromLocalStorage deletes a message from the local storage.
func (f *FuncRouter) DeleteMessageFromLocalStorage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().DeleteMessageFromLocalStorage, args...)
}

// DeleteMessage deletes a message from the server and local storage.
func (f *FuncRouter) DeleteMessage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().DeleteMessage, args...)
}

// HideAllConversations hides all conversations from the conversation list.
func (f *FuncRouter) HideAllConversations(operationID string) {
	f.call(operationID, f.userForSDK.Conversation().HideAllConversations)
}

// DeleteAllMsgFromLocalAndSvr deletes all messages from both the local storage and the server.
func (f *FuncRouter) DeleteAllMsgFromLocalAndSvr(operationID string) {
	f.call(operationID, f.userForSDK.Conversation().DeleteAllMsgFromLocalAndSvr)
}

// DeleteAllMsgFromLocal deletes all messages from the local storage only.
func (f *FuncRouter) DeleteAllMsgFromLocal(operationID string) {
	f.call(operationID, f.userForSDK.Conversation().DeleteAllMessageFromLocalStorage)
}

// ClearConversationAndDeleteAllMsg clears a conversation and deletes all associated messages.
func (f *FuncRouter) ClearConversationAndDeleteAllMsg(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().ClearConversationAndDeleteAllMsg, args...)
}

// DeleteConversationAndDeleteAllMsg deletes a conversation and all messages within it.
func (f *FuncRouter) DeleteConversationAndDeleteAllMsg(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().DeleteConversationAndDeleteAllMsg, args...)
}

// InsertSingleMessageToLocalStorage inserts a single message into the local storage.
func (f *FuncRouter) InsertSingleMessageToLocalStorage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().InsertSingleMessageToLocalStorage, args...)
}

// InsertGroupMessageToLocalStorage inserts a message into a group conversation in the local storage.
func (f *FuncRouter) InsertGroupMessageToLocalStorage(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().InsertGroupMessageToLocalStorage, args...)
}

// SearchLocalMessages searches for messages in the local storage.
func (f *FuncRouter) SearchLocalMessages(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SearchLocalMessages, args...)
}

// SetMessageLocalEx sets local extension data for a message.
func (f *FuncRouter) SetMessageLocalEx(operationID string, args ...any) {
	f.call(operationID, f.userForSDK.Conversation().SetMessageLocalEx, args...)
}
