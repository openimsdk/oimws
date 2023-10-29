package tjson

import (
	"errors"
	"github.com/xuexihuang/new_gonet/common"
)

type Login struct {
	UserName string
	PassWord string
}

type Processor struct {
}

func NewProcessor() *Processor {
	p := new(Processor)
	return p
}

func (p *Processor) UsePacketMode() bool {
	return false
}
func (p *Processor) Marshal(msg interface{}) (*common.TWSData, error) {
	////////////////////////////////////////////////////////////////////////////
	tsend := msg.(*common.TWSData)
	if tsend.MsgType != common.TextMsg && tsend.MsgType != common.BinaryMsg {
		return nil, errors.New("msg is not correct")
	}
	return tsend, nil
}
func (p *Processor) Unmarshal(data []byte) (interface{}, error) {
	return &Login{UserName: "nihao", PassWord: "huanglin"}, nil
}
func (p *Processor) Route(msg interface{}, userData interface{}) error {
	return nil
}
func (p *Processor) UnmarshalMul(nType int, data []byte) (interface{}, error) {
	ret := &common.TWSData{}
	if nType == common.TextMessage {
		ret.MsgType = common.TextMsg
		ret.Msg = data
	} else {
		ret.MsgType = common.BinaryMsg
		ret.Msg = data
	}
	return ret, nil
}
