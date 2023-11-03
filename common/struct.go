package common

type TAppParam struct {
	ModuleType    string
	PairSessionId string
	Robot3dId     int
	UE_Ip         string
	UE_Port       int
	UE_h5_weith   int
	UE_h5_higth   int
}
type TAgentUserData struct {
	SessionID string
	CookieVal string
	AppString string
	ProxyBody interface{}
}
type TWSData struct {
	MsgType int
	Msg     []byte
}

const (
	// MessageText is for UTF-8 encoded text messages like JSON.
	MessageText = iota + 1
	// MessageBinary is for binary messages like protobufs.
	MessageBinary
	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = 8

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = 9

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = 10
)
