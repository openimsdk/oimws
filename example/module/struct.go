package module

const (
	RESP_OP_TYPE      = "response"
	MQ_MSG_TYPE       = "mqMessage"
	HEART_CONFIG_TYPE = "heartConfig"
	CONN_CMD          = "connect"
	SUB_CMD           = "subscribe"
	UNSUB_CMD         = "unsubscribe"
	HEART_CMD         = "heart"
)

type ResponseSt struct {
	Type         string   `json:"type"`    //"response" or "mqMessage" or "heartConfig"
	Cmd          string   `json:"cmd"`     //"connect" "subscribe" "unsubscribe"
	Success      bool     `json:"success"` //
	ErrMsg       string   `json:"errMsg"`
	UserId       []string `json:"userIds"`
	Duration     int64    `json:"duration"` // progress run time ,seconds
	RequestId    string   `json:"requestId"`
	Topic        string   `json:"topic"`
	Extra        string   `json:"extra"`
	MsgTimeStamp int64    `json:"msgTimeStamp"`
	MsgSeqId     int64    `json:"msgSeqId"`
	Data         string   `json:"data"`
	Rate         int64    `json:"rate"`
}

type RequestSt struct {
	Cmd          string `json:"cmd"` // "subscribe" "unsubscribe" or "heart"
	RequestId    string `json:"requestId"`
	Topic        string `json:"topic"`
	Extra        string `json:"extra"`
	MsgTimeStamp int64  `json:"msgStartTime"`
	MsgSeqId     int64  `json:"msgSeqId"`
}
