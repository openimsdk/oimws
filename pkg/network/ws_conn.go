package network

import (
	"errors"
	"github.com/openim-sigs/oimws/pkg/common"
	"net"
	"sync"

	"github.com/gorilla/websocket"
	log "github.com/xuexihuang/new_log15"
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

// newWSConn initializes a new WSConn object.
func newWSConn(conn *websocket.Conn, pendingWriteNum int, maxMsgLen uint32, appurl string, cookieVal string) *WSConn {
	log.Error("test4.1", pendingWriteNum)
	wsConn := new(WSConn)
	wsConn.conn = conn
	wsConn.writeChan = make(chan *common.TWSData, pendingWriteNum)
	log.Error("test4.1.1", pendingWriteNum)
	wsConn.maxMsgLen = maxMsgLen
	/////////////////////////////生成唯一sessionid。
	var sessionID string
	log.Error("test4.1.2", pendingWriteNum)
	//sessionID = common.GetRandomSessionId()

	log.Error("test4.1.3", pendingWriteNum)
	wsConn.SessionId = sessionID
	wsConn.AppURL = appurl
	wsConn.CookieVal = cookieVal
	log.Error("test4.2")
	go func() {
		for b := range wsConn.writeChan {
			if b == nil {
				break
			}
			var err error
			if b.MsgType == common.MessageBinary {
				err = conn.WriteMessage(websocket.BinaryMessage, b.Msg)
			} else if b.MsgType == common.MessageText {
				err = conn.WriteMessage(websocket.TextMessage, b.Msg)
			} else if b.MsgType == common.PingMessage {
				log.Info("ping message", "b", b)
				err = conn.WriteMessage(websocket.PingMessage, b.Msg)
			} else if b.MsgType == common.CloseMessage {
				log.Info("close message", "b", b)
				err = conn.WriteMessage(websocket.CloseMessage, b.Msg)
				break
			}
			if err != nil {
				log.Error("send message err", "err", err.Error())
				break
			}
			//fmt.Println("send msg is :", b)
		}

		conn.Close()
		wsConn.Lock()
		wsConn.closeFlag = true
		wsConn.Unlock()
	}()
	log.Error("test4.3")
	return wsConn
}

// doDestroy forcefully closes the connection without waiting for pending writes.
func (wsConn *WSConn) doDestroy() {
	wsConn.conn.UnderlyingConn().(*net.TCPConn).SetLinger(0)
	wsConn.conn.Close()

	if !wsConn.closeFlag {
		close(wsConn.writeChan)
		wsConn.closeFlag = true
	}
}

// Destroy cleanly closes the connection.
func (wsConn *WSConn) Destroy() {
	wsConn.Lock()
	defer wsConn.Unlock()

	wsConn.doDestroy()
}

// Close initiates a graceful shutdown of the connection.
func (wsConn *WSConn) Close() {
	wsConn.Lock()
	defer wsConn.Unlock()
	if wsConn.closeFlag {
		return
	}

	wsConn.doWrite(nil)
	wsConn.closeFlag = true
}

// doWrite enqueues a message for writing to the websocket connection.
func (wsConn *WSConn) doWrite(b *common.TWSData) {
	if len(wsConn.writeChan) == cap(wsConn.writeChan) {
		//log.Debug("close conn: channel full")
		log.Error("close conn: channel full")
		wsConn.doDestroy()
		return
	}

	wsConn.writeChan <- b
}

// LocalAddr returns the local network address.
func (wsConn *WSConn) LocalAddr() net.Addr {
	return wsConn.conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (wsConn *WSConn) RemoteAddr() net.Addr {
	return wsConn.conn.RemoteAddr()
}

// ReadMsg reads a message from the websocket connection.(goroutine not safe)
// goroutine not safe.
func (wsConn *WSConn) ReadMsg() (int, []byte, error) {
	nTye, b, err := wsConn.conn.ReadMessage()
	return nTye, b, err
}

// WriteMsg writes a message to the websocket connection.(Args must not be modified by the others goroutines)
// args must not be modified by the others goroutines.
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
	}

	wsConn.doWrite(args)
	return nil

}
