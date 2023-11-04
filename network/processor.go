package network

import "github.com/openim-sigs/oimws/common"

type Processor interface {
	// must goroutine safe
	Route(msg interface{}, userData interface{}) error
	// must goroutine safe
	Unmarshal(data []byte) (interface{}, error)
	UnmarshalMul(nType int, data []byte) (interface{}, error)
	// must goroutine safe
	Marshal(msg interface{}) (*common.TWSData, error)
	// 是否使用压解包模式
	UsePacketMode() bool
}
