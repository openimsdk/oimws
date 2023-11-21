# OIMWS - OpenIM WebSocket Service 😎

![Build Status](https://github.com/openim-sigs/oimws/actions/workflows/test.yml/badge.svg)
[![Codecov](https://img.shields.io/codecov/c/github/openim-sigs/oimws)](https://app.codecov.io/github/openim-sigs/oimws)
[![Go Report Card](https://goreportcard.com/badge/github.com/openim-sigs/oimws)](https://goreportcard.com/report/github.com/openim-sigs/oimws)
[![GoDoc](https://godoc.org/github.com/openim-sigs/oimws?status.svg)](https://godoc.org/github.com/openim-sigs/oimws)

> :notes: OIMWS (OpenIM WebSocket Service) is a high-performance, scalable WebSocket framework designed specifically for building instant messaging (IM) systems. Harnessing the concurrent capabilities of Go and the real-time communication provided by WebSocket protocol, OIMWS offers a robust backend solution to support modern instant communication needs, ranging from basic message transit to complex session management and network optimization.

## Features 🚀

+ **WebSocket Support**: Provides full WebSocket connection handling for high-concurrency client communication.
+ **Modular Architecture**: Features a set of decoupled modules for message processing, connection management, and user/group interactions, facilitating development and maintenance.
+ **Real-Time Message Processing**: Ensures swift response and distribution of real-time messages for timely and reliable communication.
+ **Network Layer Abstraction**: Includes low-level network abstractions allowing for customized protocol and data process flows.
+ **Configuration Management**: Simplified configuration management supports rapid deployment and environmental adaptation.
+ **Utility Toolkit**: Offers a wide array of utility functions for common operations such as date processing and text manipulation.
+ **Automated Testing**: Built-in unit tests ensure the stability of modules and features.
+ **Clear Build Scripts**: Makefile support simplifies the build process, enhancing developer productivity.
+ **Code Quality Assurance**: Integrates `golangci-lint` to ensure consistency in code quality and style.
+ **Sample Code**: Includes examples to demonstrate framework usage, use openim jssdk(10003) accelerating the learning curve for developers.

## Quick Start 🚗💨

Clone the repository to your local machine:
```bash
git clone https://github.com/openim-sigs/oimws.git
cd oimws
```

**Build oimws**

```bash
make build
```

**Start oimws**

```bash
make start
```


### [Example](https://github.com/openim-sigs/oimws/tree/main/example)

An encapsulated framework within `jssdk` connecting to `openim-sdk-core`, providing streamlined management and integration of WebSocket, TCP, and HTTP protocols in the OpenIM ecosystem.


### Usage

Import the necessary modules and initialize the protocol framework:

```go
func main() {
	var sdkWsPort, logLevel *int
	var openIMWsAddress, openIMApiAddress, openIMDbDir *string
	openIMApiAddress = flag.String("openIM_api_address", "http://127.0.0.1:10002",
		"openIM api listening address")
	openIMWsAddress = flag.String("openIM_ws_address", "http://127.0.0.1:10001",
		"openIM ws listening address")
	sdkWsPort = flag.Int("sdk_ws_port", 10003, "openIMSDK ws listening port")
	logLevel = flag.Int("openIM_log_level", 6, "control log output level")
	openIMDbDir = flag.String("openIMDbDir", "./", "openIM db dir")
	flag.Parse()
	core_func.Config.WsAddr = *openIMWsAddress
	core_func.Config.ApiAddr = *openIMApiAddress
	core_func.Config.DataDir = *openIMDbDir
	core_func.Config.LogLevel = uint32(*logLevel)
	core_func.Config.IsLogStandardOutput = true
	log.SetOutLevel(log.LvlInfo)
	fmt.Println("Client starting....")
	log.Info("Client starting....")
	gatenet := Initsever(*sdkWsPort)
	gatenet.SetMsgFun(module.NewAgent, module.CloseAgent, module.DataRecv)
	go gatenet.Runloop()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGTERM)
	sig := <-c
	log.Info("wsconn server closing down ", "sig", sig)
	gatenet.CloseGate()
}
```

### [More examples](https://github.com/openim-sigs/oimws/tree/master/examples)


## Contribution Ⓜ️

Feel free to contribute to this project by opening issues or submitting pull requests.

<a href="https://github.com/openim-sigs/oimws/graphs/contributors">
	<img src="https://contrib.rocks/image?repo=openim-sigs/oimws" />
</a>

## License 🤝

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
