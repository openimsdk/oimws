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
	MsgType WsMesgType
	Msg     []byte
}
type WsMesgType int

const (
	TextMsg     WsMesgType = 1
	BinaryMsg   WsMesgType = 0
	CloseMsg    WsMesgType = 2
	TextMessage int        = 1
)
