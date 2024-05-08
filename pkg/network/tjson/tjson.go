package tjson

import (
	"github.com/openim-sigs/oimws/pkg/common"
)

type Login struct {
	UserName string
	PassWord string
}

type Processor struct {
}

// NewProcessor is a constructor for Processor.
func NewProcessor() *Processor {
	p := new(Processor)
	return p
}

// UsePacketMode returns false indicating that the processor is likely used in a stream mode and not packet mode.
func (p *Processor) UsePacketMode() bool {
	return false
}

// Marshal takes a message interface and converts it into a WebSocket data structure.
func (p *Processor) Marshal(msg interface{}) (*common.TWSData, error) {
	tsend := msg.(*common.TWSData)
	//if tsend.MsgType != common.MessageText && tsend.MsgType != common.MessageBinary {
	//	return nil, errors.New("msg is not correct")
	//}
	return tsend, nil
}

// Unmarshal creates a new Login struct with preset credentials, not actually using the input data.
func (p *Processor) Unmarshal(data []byte) (interface{}, error) {
	return &Login{UserName: "nihao", PassWord: "huanglin"}, nil
}

// Route currently does nothing and always returns nil, indicating no error.
func (p *Processor) Route(msg interface{}, userData interface{}) error {
	return nil
}

// UnmarshalMul takes a message type and data, wraps it into a TWSData struct, and returns it.
func (p *Processor) UnmarshalMul(nType int, data []byte) (interface{}, error) {
	ret := &common.TWSData{}
	if nType == common.MessageText {
		ret.MsgType = common.MessageText
		ret.Msg = data
	} else {
		ret.MsgType = common.MessageBinary
		ret.Msg = data
	}
	return ret, nil
}
