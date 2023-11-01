# OpenIMProtocolFramework

![Build Status](https://github.com/openim-sigs/oimws/actions/workflows/test.yml/badge.svg)
[![Codecov](https://img.shields.io/codecov/c/github/openim-sigs/oimws)](https://app.codecov.io/github/openim-sigs/oimws)
[![Go Report Card](https://goreportcard.com/badge/github.com/openim-sigs/oimws)](https://goreportcard.com/report/github.com/openim-sigs/oimws)
[![GoDoc](https://godoc.org/github.com/openim-sigs/oimws?status.svg)](https://godoc.org/github.com/openim-sigs/oimws)

> :notes: Minimalist websocket framework for Go.

Melody is websocket framework based on [github.com/gorilla/websocket](https://github.com/gorilla/websocket)
that abstracts away the tedious parts of handling websockets. It gets out of
your way so you can write real-time apps. Features include:

* [x] Clear and easy interface similar to `net/http` or Gin.
* [x] A simple way to broadcast to all or selected connected sessions.
* [x] Message buffers making concurrent writing safe.
* [x] Automatic handling of sending ping/pong heartbeats that timeout broken sessions.
* [x] Store data on sessions.

## Install

```bash
go get github.com/openim-sigs/oimws
```

## [Example: chat](https://github.com/openim-sigs/oimws/tree/master/examples/chat)

An encapsulated framework within `jssdk` connecting to `openim-sdk-core`, providing streamlined management and integration of WebSocket, TCP, and HTTP protocols in the OpenIM ecosystem.

## Features

- **Protocol Management**: Seamless management and integration of WebSocket, TCP, and HTTP protocols.
- **Encoding and Decoding**: Handle message encoding and decoding for cross-language compatibility.
- **Protocol Implementation**: Implement basic OpenIM protocols like login, push, etc.
- **Event Handling**: Convert received messages into corresponding events and pass them to upper layer applications for processing.

## Getting Started

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/openim-sigs/oimws.git
cd oimws
```


### Usage

Import the necessary modules and initialize the protocol framework:

```go
go
```

### [More examples](https://github.com/openim-sigs/oimws/tree/master/examples)



## Contribution

Feel free to contribute to this project by opening issues or submitting pull requests.

<a href="https://github.com/openim-sigs/oimws/graphs/contributors">
	<img src="https://contrib.rocks/image?repo=openim-sigs/oimws" />
</a>

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
