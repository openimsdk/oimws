package network

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/xuexihuang/new_gonet/common"
	log "github.com/xuexihuang/new_log15"
	"net"
	"sync"
)

type WebsocketConnSet map[*websocket.Conn]struct{}

type WSConn struct {
	sync.Mutex
	conn      *websocket.Conn
	writeChan chan *common.TWSData
	maxMsgLen uint32
	closeFlag bool
	//add by hl
	SessionId string
	AppParam  common.TAppParam
	AppURL    string
	CookieVal string
}

func newWSConn(conn *websocket.Conn, pendingWriteNum int, maxMsgLen uint32, appurl string, cookieVal string) *WSConn {
	wsConn := new(WSConn)
	wsConn.conn = conn
	wsConn.writeChan = make(chan *common.TWSData, pendingWriteNum)
	wsConn.maxMsgLen = maxMsgLen
	/////////////////////////////生成唯一sessionid。
	var sessionID string
	sessionID = common.GetRandomSessionId()
	wsConn.SessionId = sessionID
	wsConn.AppURL = appurl
	wsConn.CookieVal = cookieVal

	go func() {
		for b := range wsConn.writeChan {
			if b == nil {
				break
			}
			var err error
			if b.MsgType == common.BinaryMsg {
				err = conn.WriteMessage(websocket.BinaryMessage, b.Msg)
			} else if b.MsgType == common.TextMsg {
				err = conn.WriteMessage(websocket.TextMessage, b.Msg)
			}
			if err != nil {
				break
			}
			//fmt.Println("send msg is :", b)
		}

		conn.Close()
		wsConn.Lock()
		wsConn.closeFlag = true
		wsConn.Unlock()
	}()

	return wsConn
}

func (wsConn *WSConn) doDestroy() {
	wsConn.conn.UnderlyingConn().(*net.TCPConn).SetLinger(0)
	wsConn.conn.Close()

	if !wsConn.closeFlag {
		close(wsConn.writeChan)
		wsConn.closeFlag = true
	}
}

func (wsConn *WSConn) Destroy() {
	wsConn.Lock()
	defer wsConn.Unlock()

	wsConn.doDestroy()
}

func (wsConn *WSConn) Close() {
	wsConn.Lock()
	defer wsConn.Unlock()
	if wsConn.closeFlag {
		return
	}

	wsConn.doWrite(nil)
	wsConn.closeFlag = true
}

func (wsConn *WSConn) doWrite(b *common.TWSData) {
	if len(wsConn.writeChan) == cap(wsConn.writeChan) {
		//log.Debug("close conn: channel full")
		log.Error("close conn: channel full")
		wsConn.doDestroy()
		return
	}

	wsConn.writeChan <- b
}

func (wsConn *WSConn) LocalAddr() net.Addr {
	return wsConn.conn.LocalAddr()
}

func (wsConn *WSConn) RemoteAddr() net.Addr {
	return wsConn.conn.RemoteAddr()
}

// goroutine not safe
func (wsConn *WSConn) ReadMsg() (int, []byte, error) {
	nTye, b, err := wsConn.conn.ReadMessage()
	return nTye, b, err
}

// args must not be modified by the others goroutines
func (wsConn *WSConn) WriteMsg(args *common.TWSData) error {
	wsConn.Lock()
	defer wsConn.Unlock()
	if wsConn.closeFlag {
		return nil
	}

	// get len
	var msgLen uint32
	msgLen = uint32(len(args.Msg))

	// check len
	if msgLen > wsConn.maxMsgLen {
		return errors.New("message too long")
	} else if msgLen < 1 {
		return errors.New("message too short")
	}

	wsConn.doWrite(args)
	return nil

}
