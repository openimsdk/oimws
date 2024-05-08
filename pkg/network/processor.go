package network

import (
	"github.com/openim-sigs/oimws/pkg/common"
)

type Processor interface {
	// must goroutine safe
	Route(msg interface{}, userData interface{}) error
	// must goroutine safe
	Unmarshal(data []byte) (interface{}, error)
	UnmarshalMul(nType int, data []byte) (interface{}, error)
	// must goroutine safe
	Marshal(msg interface{}) (*common.TWSData, error)
	// Whether to use packet mode for packing/unpacking
	UsePacketMode() bool
}
