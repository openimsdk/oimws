package gate

import (
	common2 "github.com/openim-sigs/oimws/pkg/common"
	network2 "github.com/openim-sigs/oimws/pkg/network"
	"net"
	"reflect"
	"time"

	log "github.com/xuexihuang/new_log15"
)

type Gate struct {
	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	Processor       network2.Processor
	//AgentChanRPC    *chanrpc.Server

	// websocket
	WSAddr      string
	HTTPTimeout time.Duration
	CertFile    string
	KeyFile     string

	// tcp
	TCPAddr   string
	LenMsgLen int

	//add by huanglin
	FunNewAgent   func(Agent)
	FunCloseAgent func(Agent)
	FuncMsgRecv   func(interface{}, Agent)
}

func NewGate(maxConnNum int, maxMsgLen uint32, processor network2.Processor, WSAddr string,
	HTTPTimeout time.Duration, writerChanLen int) *Gate {
	return &Gate{MaxConnNum: maxConnNum, MaxMsgLen: maxMsgLen, Processor: processor, WSAddr: WSAddr,
		HTTPTimeout: HTTPTimeout, PendingWriteNum: writerChanLen}
}

// SetFun sets the functions for handling new agents, closing agents, and receiving messages.
func (gate *Gate) SetFun(Fun1 func(Agent), Fun2 func(Agent), Fun3 func(interface{}, Agent)) {
	gate.FunNewAgent = Fun1
	gate.FunCloseAgent = Fun2
	gate.FuncMsgRecv = Fun3
}

// Run starts the gate service and listens for incoming connections.
func (gate *Gate) Run(closeSig chan bool) {
	var wsServer *network2.WSServer
	if gate.WSAddr != "" {
		wsServer = new(network2.WSServer)
		wsServer.Addr = gate.WSAddr
		wsServer.MaxConnNum = gate.MaxConnNum
		wsServer.PendingWriteNum = gate.PendingWriteNum
		wsServer.MaxMsgLen = gate.MaxMsgLen
		wsServer.HTTPTimeout = gate.HTTPTimeout
		wsServer.CertFile = gate.CertFile
		wsServer.KeyFile = gate.KeyFile
		wsServer.NewAgent = func(conn *network2.WSConn) network2.Agent {
			a := &agent{conn: conn, gate: gate}
			/*if gate.AgentChanRPC != nil {
				gate.AgentChanRPC.Go("NewAgent", a)
			}*/
			/////////////////////////////////////////////////////
			tagent := common2.TAgentUserData{SessionID: conn.SessionId, AppString: conn.AppURL, CookieVal: conn.CookieVal}
			a.SetUserData(&tagent)
			gate.FunNewAgent(a)
			return a
		}
	}
	/*
		var tcpServer *network.TCPServer
		if gate.TCPAddr != "" {
			tcpServer = new(network.TCPServer)
			tcpServer.Addr = gate.TCPAddr
			tcpServer.MaxConnNum = gate.MaxConnNum
			tcpServer.PendingWriteNum = gate.PendingWriteNum
			tcpServer.LenMsgLen = gate.LenMsgLen
			tcpServer.MaxMsgLen = gate.MaxMsgLen
			tcpServer.UsePacketMode = gate.Processor.UsePacketMode()
			tcpServer.NewAgent = func(conn *network.TCPConn) network.Agent {
				a := &agent{conn: conn, gate: gate}
				if gate.AgentChanRPC != nil {
					gate.AgentChanRPC.Go("NewAgent", a)
				}
				gate.FunNewAgent(a)
				return a
			}
		}*/

	if wsServer != nil {
		wsServer.Start()
	}
	/*if tcpServer != nil {
		tcpServer.Start()
	}*/
	<-closeSig
	if wsServer != nil {
		wsServer.Close()
	}
	/*if tcpServer != nil {
		tcpServer.Close()
	}*/
}

func (gate *Gate) OnDestroy() {}

type agent struct {
	conn     network2.Conn
	gate     *Gate
	userData interface{}
}

// Run processes incoming messages in a loop.
func (a *agent) Run() {
	defer common2.TryRecoverAndDebugPrint()
	for {
		nType, data, err := a.conn.ReadMsg()
		if err != nil {
			//log.Debug("read message: %v", err)
			log.Info("read message error", "error", err)
			break
		}
		log.Debug("recve one ws msg ", "nType", nType)
		if a.gate.Processor != nil {
			msg, err := a.gate.Processor.UnmarshalMul(nType, data)
			if err != nil {
				//log.Debug("unmarshal message error: %v", err)
				log.Error("unmarshal message error", "err", err)
				break
			}
			a.gate.FuncMsgRecv(msg, a)
			/*err = a.gate.Processor.Route(msg, a)
			if err != nil {
				log.Debug("route message error: %v", err)
				break
			}*/
		}
	}
}

// OnClose is called when the agent's connection is closed.
func (a *agent) OnClose() {

	/*if a.gate.AgentChanRPC != nil {
		err := a.gate.AgentChanRPC.Call0("CloseAgent", a)
		if err != nil {
			log.Error("chanrpc error: %v", err)
		}
	}*/
	a.gate.FunCloseAgent(a)
}

// WriteMsg sends a message to the client.
func (a *agent) WriteMsg(msg interface{}) {
	if a.gate.Processor != nil {
		data, err := a.gate.Processor.Marshal(msg)
		if err != nil {
			//log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
			log.Error("marshal message", "reflect.TypeOf(msg)", reflect.TypeOf(msg), "error", err)
			return
		}
		err = a.conn.WriteMsg(data)
		if err != nil {
			//log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
			log.Error("write message error", "reflect.TypeOf(msg)", reflect.TypeOf(msg), "error", err)
		}
	}
}

func (a *agent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *agent) Close() {
	a.conn.Close()
}

func (a *agent) Destroy() {
	a.conn.Destroy()
}

func (a *agent) UserData() interface{} {
	return a.userData
}

func (a *agent) SetUserData(data interface{}) {
	a.userData = data
}
