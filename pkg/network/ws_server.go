package network

import (
	"crypto/tls"
	"fmt"
	"github.com/openim-sigs/oimws/pkg/common"
	"net"
	"net/http"
	"sync"
	"time"

	log "github.com/xuexihuang/new_log15"

	"github.com/gorilla/websocket"
)

type WSServer struct {
	Addr            string
	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	HTTPTimeout     time.Duration
	CertFile        string
	KeyFile         string
	NewAgent        func(*WSConn) Agent
	ln              net.Listener
	handler         *WSHandler
}

type WSHandler struct {
	maxConnNum      int
	pendingWriteNum int
	maxMsgLen       uint32
	newAgent        func(*WSConn) Agent
	upgrader        websocket.Upgrader
	conns           WebsocketConnSet
	mutexConns      sync.Mutex
	wg              sync.WaitGroup
}

// ServeHTTP handles HTTP requests and upgrades them to WebSocket if the request is valid.
func (handler *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer common.TryRecoverAndDebugPrint()
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	var cookieVal string
	/*
		cookieToken, err := r.Cookie("token")
		if err != nil {
			log.Info("http的cookie中没有对应token", "err", err)
		} else {
			cookieVal = cookieToken.Value
		}*/
	token := r.Header.Get("Authorization")
	if token != "" && len(token) > 7 {
		cookieVal = token[7:]
	} else {
		log.Info("http的headers Authorization中没有对应token")
	}
	log.Info("token info", "token", cookieVal)

	fmt.Println("url is:", r.URL.Path)
	conn, err := handler.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("upgrade error", "err", err, "remoteIp", r.Host)
		return
	}
	conn.SetReadLimit(int64(handler.maxMsgLen))
	_ = conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	conn.SetPongHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		log.Info("js replying with a pong packet.")
		return nil
	})
	log.Error("test1")
	handler.wg.Add(1)
	defer handler.wg.Done()

	handler.mutexConns.Lock()
	if handler.conns == nil {
		handler.mutexConns.Unlock()
		conn.Close()
		return
	}
	log.Error("test2")
	if len(handler.conns) >= handler.maxConnNum {
		handler.mutexConns.Unlock()
		conn.Close()
		log.Error("too many connections")
		return
	}
	log.Error("test3")
	handler.conns[conn] = struct{}{}
	handler.mutexConns.Unlock()

	log.Error("test4")
	wsConn := newWSConn(conn, handler.pendingWriteNum, handler.maxMsgLen, r.URL.String(), cookieVal)
	log.Error("tes5")
	agent := handler.newAgent(wsConn)
	agent.Run()

	// cleanup
	wsConn.Close()
	handler.mutexConns.Lock()
	delete(handler.conns, conn)
	handler.mutexConns.Unlock()
	agent.OnClose()
}

// Start initializes and starts the WebSocket server.
func (server *WSServer) Start() {
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		//log.Fatal("%v", err)
		log.Crit("net.listen err", "err", err)
	}

	if server.MaxConnNum <= 0 {
		server.MaxConnNum = 100
		//log.Release("invalid MaxConnNum, reset to %v", server.MaxConnNum)
		log.Info("invalid MaxConnNum,reset", "server.MaxConnNum", server.MaxConnNum)
	}
	if server.PendingWriteNum <= 0 {
		server.PendingWriteNum = 100
		//log.Release("invalid PendingWriteNum, reset to %v", server.PendingWriteNum)
		log.Info("invalid PendingWriteNum,reset", "server.PendingWriteNum", server.PendingWriteNum)
	}
	if server.MaxMsgLen <= 0 {
		server.MaxMsgLen = 4096
		//log.Release("invalid MaxMsgLen, reset to %v", server.MaxMsgLen)
		log.Info("invalid MaxMsgLen,reset", "server.MaxMsgLen", server.MaxMsgLen)
	}
	if server.HTTPTimeout <= 0 {
		server.HTTPTimeout = 10 * time.Second
		//log.Release("invalid HTTPTimeout, reset to %v", server.HTTPTimeout)
		log.Info("invalid HTTPTimeout,reset", "server.HTTPTimeout", server.HTTPTimeout)
	}
	if server.NewAgent == nil {
		//log.Fatal("NewAgent must not be nil")
		log.Crit("NewAgent must not be nil")
	}

	if server.CertFile != "" || server.KeyFile != "" {
		config := &tls.Config{}
		config.NextProtos = []string{"http/1.1"}

		var err error
		config.Certificates = make([]tls.Certificate, 1)
		config.Certificates[0], err = tls.LoadX509KeyPair(server.CertFile, server.KeyFile)
		if err != nil {
			//log.Fatal("%v", err)
			log.Crit("cerfiti file error", "err", err)
		}

		ln = tls.NewListener(ln, config)
	}

	server.ln = ln
	server.handler = &WSHandler{
		maxConnNum:      server.MaxConnNum,
		pendingWriteNum: server.PendingWriteNum,
		maxMsgLen:       server.MaxMsgLen,
		newAgent:        server.NewAgent,
		conns:           make(WebsocketConnSet),
		upgrader: websocket.Upgrader{
			HandshakeTimeout: server.HTTPTimeout,
			CheckOrigin:      func(_ *http.Request) bool { return true },
		},
	}

	httpServer := &http.Server{
		Addr:           server.Addr,
		Handler:        server.handler,
		ReadTimeout:    server.HTTPTimeout,
		WriteTimeout:   server.HTTPTimeout,
		MaxHeaderBytes: 1024,
	}

	go httpServer.Serve(ln)
}

// Close shuts down the WebSocket server and closes all active connections.
func (server *WSServer) Close() {
	server.ln.Close()

	server.handler.mutexConns.Lock()
	for conn := range server.handler.conns {
		conn.Close()
	}
	server.handler.conns = nil
	server.handler.mutexConns.Unlock()

	server.handler.wg.Wait()
}
