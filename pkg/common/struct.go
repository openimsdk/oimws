package common

// TAppParam defines the configuration parameters related to an application.
type TAppParam struct {
	ModuleType    string // The type of the module, used to identify different modules or features.
	PairSessionId string // The unique identifier for a pairing session.
	Robot3dId     int    // The unique identifier for a 3D robot.
	UE_Ip         string // The IP address of the User Equipment (UE).
	UE_Port       int    // The port number of the User Equipment (UE).
	UE_h5_weith   int    // The width of the HTML5 element on the User Equipment (UE), likely a typo and should be 'UE_h5_width'.
	UE_h5_higth   int    // The height of the HTML5 element on the User Equipment (UE), likely a typo and should be 'UE_h5_height'.
}

// TAgentUserData contains user-specific data passed to an agent.
type TAgentUserData struct {
	SessionID string      // The session ID uniquely identifying the user session.
	CookieVal string      // The value of the cookie stored for the user.
	AppString string      // A string related to the application, possibly containing user-specific settings or states.
	ProxyBody interface{} // A generic interface to hold different types of data for proxy communication.
	UserId    string
}

// TWSData defines the structure for WebSocket data transmission.
type TWSData struct {
	MsgType int    // The type of the message, used to handle different data or requests.
	Msg     []byte // The actual message data in bytes.
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
