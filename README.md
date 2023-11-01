# OpenIMProtocolFramework

![Build Status](https://github.com/olahol/melody/actions/workflows/test.yml/badge.svg)
[![Codecov](https://img.shields.io/codecov/c/github/olahol/melody)](https://app.codecov.io/github/olahol/melody)
[![Go Report Card](https://goreportcard.com/badge/github.com/olahol/melody)](https://goreportcard.com/report/github.com/olahol/melody)
[![GoDoc](https://godoc.org/github.com/olahol/melody?status.svg)](https://godoc.org/github.com/olahol/melody)

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
go get github.com/olahol/melody
```

## [Example: chat](https://github.com/olahol/melody/tree/master/examples/chat)

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
git clone https://github.com/your-username/oimws.git
cd OpenIMProtocolFramework
```


### Usage

Import the necessary modules and initialize the protocol framework:

```go
go
```

### [More examples](https://github.com/olahol/melody/tree/master/examples)



## Contribution

Feel free to contribute to this project by opening issues or submitting pull requests.

<a href="https://github.com/olahol/melody/graphs/contributors">
	<img src="https://contrib.rocks/image?repo=olahol/melody" />
</a>

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
