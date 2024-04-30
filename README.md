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
mage
```

**Start oimws**

```bash
mage start
```

**Check oimws status**

```bash
mage check
```

**Stop oimws Status:**

```bash
mage stop
```

### [main](https://github.com/openim-sigs/oimws/tree/main/cmd)

An encapsulated framework within `jssdk` connecting to `openim-sdk-core`, providing streamlined management and integration of WebSocket, TCP, and HTTP protocols in the OpenIM ecosystem.


### code

the folders of oimws:

cmd--------------  the main.go folder


common-----------  common structures and functions, primarily used by the network frame


core_func--------  Some functions encapsulate the interface for calling the JS SDK.


gate-------------  network frame,functions for websocket


module-----------  the module codes


network----------  network frame



### Cluster solution(Consistent Hashing)

Using consistent hashing in Nginx typically involves the hash directive within the upstream module. 
Starting from Nginx 1.7.2, it supports consistent hashing based on specified variables. 
You can use the request parameter sendId as the key for consistent hashing to distribute requests.

Here's a configuration example that demonstrates how to use consistent hashing for the /ws endpoint to select backend servers:
```
http {
    upstream backend {
        # Use userId as the key for consistent hashing
        hash $arg_sendId consistent;

        # Define backend servers
        server backend1.example.com;
        server backend2.example.com;
        # ... More backend servers ...
    }

    server {
        listen 80;

        location /ws {
            proxy_pass http://backend;
            # ... Other possible proxy settings ...
        }

        # ... Other location definitions ...
    }
}

```
## Contribution Ⓜ️

Feel free to contribute to this project by opening issues or submitting pull requests.

<a href="https://github.com/openim-sigs/oimws/graphs/contributors">
	<img src="https://contrib.rocks/image?repo=openim-sigs/oimws" />
</a>

## License 🤝

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
