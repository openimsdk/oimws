package network

import (
	"net"

	"github.com/openim-sigs/oimws/common"
)

type Conn interface {
	ReadMsg() (int, []byte, error)
	WriteMsg(args *common.TWSData) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
}
