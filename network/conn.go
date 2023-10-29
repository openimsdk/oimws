package network

import (
	"github.com/xuexihuang/new_gonet/common"
	"net"
)

type Conn interface {
	ReadMsg() (int, []byte, error)
	WriteMsg(args *common.TWSData) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
}
