package network

import (
	"github.com/openim-sigs/oimws/pkg/common"
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
