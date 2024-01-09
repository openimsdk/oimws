package network

import (
	"sync"
	"time"

	log "github.com/xuexihuang/new_log15"

	"github.com/gorilla/websocket"
)

const (
	MaxMsgLen = 1024 * 1024 * 10
)

// WSClient just for client dial to websocket server
type WSClient struct {
	sync.Mutex
	Addr             string
	ConnNum          int
	ConnectInterval  time.Duration
	PendingWriteNum  int
	MaxMsgLen        uint32
	HandshakeTimeout time.Duration
	AutoReconnect    bool
	NewAgent         func(*WSConn) Agent
	dialer           websocket.Dialer
	conns            WebsocketConnSet
	wg               sync.WaitGroup
	closeFlag        bool
}

// Start initializes and starts the WebSocket client.
func (client *WSClient) Start() {
	client.init()

	for i := 0; i < client.ConnNum; i++ {
		client.wg.Add(1)
		go client.connect()
	}
}

// init prepares the client by setting default values and validating settings.
func (client *WSClient) init() {
	client.Lock()
	defer client.Unlock()

	if client.ConnNum <= 0 {
		client.ConnNum = 1
		//log.Release("invalid ConnNum, reset to %v", client.ConnNum)
		log.Info("invalid ConnNum, reset", "client.ConnNum", client.ConnNum)
	}
	if client.ConnectInterval <= 0 {
		client.ConnectInterval = 3 * time.Second
		//log.Release("invalid ConnectInterval, reset to %v", client.ConnectInterval)
		log.Info("invalid ConnectInterval, reset", "client.ConnectInterval", client.ConnectInterval)
	}
	if client.PendingWriteNum <= 0 {
		client.PendingWriteNum = 100
		//log.Release("invalid PendingWriteNum, reset to %v", client.PendingWriteNum)
		log.Info("invalid PendingWriteNum, reset", "client.PendingWriteNum", client.PendingWriteNum)
	}
	if client.MaxMsgLen <= 0 {
		client.MaxMsgLen = MaxMsgLen
		//log.Release("invalid MaxMsgLen, reset to %v", client.MaxMsgLen)
		log.Info("invalid MaxMsgLen, reset", "client.MaxMsgLen", client.MaxMsgLen)
	}
	if client.HandshakeTimeout <= 0 {
		client.HandshakeTimeout = 10 * time.Second
		//log.Release("invalid HandshakeTimeout, reset to %v", client.HandshakeTimeout)
		log.Info("invalid HandshakeTimeout, reset", "client.HandshakeTimeout", client.HandshakeTimeout)
	}
	if client.NewAgent == nil {
		//log.Fatal("NewAgent must not be nil")
		log.Crit("NewAgent must not be nil")
	}
	if client.conns != nil {
		//log.Fatal("client is running")
		log.Crit("client is running")
	}

	client.conns = make(WebsocketConnSet)
	client.closeFlag = false
	client.dialer = websocket.Dialer{
		HandshakeTimeout: client.HandshakeTimeout,
	}
}

// dial creates a new WebSocket connection.
func (client *WSClient) dial() *websocket.Conn {
	for {
		conn, _, err := client.dialer.Dial(client.Addr, nil)
		if err == nil || client.closeFlag {
			return conn
		}

		//log.Release("connect to %v error: %v", client.Addr, err)
		log.Info("connect error", "client.Addr", client.Addr, "err", err)
		time.Sleep(client.ConnectInterval)
		continue
	}
}

// connect handles the connection lifecycle.
func (client *WSClient) connect() {
	defer client.wg.Done()

reconnect:
	conn := client.dial()
	if conn == nil {
		return
	}
	conn.SetReadLimit(int64(client.MaxMsgLen))

	client.Lock()
	if client.closeFlag {
		client.Unlock()
		conn.Close()
		return
	}
	client.conns[conn] = struct{}{}
	client.Unlock()

	wsConn := newWSConn(conn, client.PendingWriteNum, client.MaxMsgLen, "", "")
	agent := client.NewAgent(wsConn)
	agent.Run()

	// cleanup
	wsConn.Close()
	client.Lock()
	delete(client.conns, conn)
	client.Unlock()
	agent.OnClose()

	if client.AutoReconnect {
		time.Sleep(client.ConnectInterval)
		goto reconnect
	}
}

// Close initiates the shutdown process for the client.
func (client *WSClient) Close() {
	client.Lock()
	client.closeFlag = true
	for conn := range client.conns {
		conn.Close()
	}
	client.conns = nil
	client.Unlock()

	client.wg.Wait()
}
