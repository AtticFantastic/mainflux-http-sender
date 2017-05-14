# Mainflux HTTP Sender

[![License](https://img.shields.io/badge/license-Apache%20v2.0-blue.svg)](LICENSE)
[![Build Status](https://travis-ci.org/mainflux/mainflux-http-sender.svg?branch=master)](https://travis-ci.org/mainflux/mainflux-http-sender)
[![Go Report Card](https://goreportcard.com/badge/github.com/mainflux/mainflux-http-sender)](https://goreportcard.com/report/github.com/mainflux/mainflux-http-sender)
[![Join the chat at https://gitter.im/Mainflux/mainflux](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Mainflux HTTP API for message sending to NATS broker (which emits them to other services).

### Installation
Use [`go`](https://golang.org/cmd/go/) tool to "get" (i.e. fetch and build) `mainflux-http-sender` package:
```bash
go get github.com/mainflux/mainflux-http-sender
```

This will download the code to `$GOPATH/src/github.com/mainflux/mainflux-http-sender` directory,
and then compile it and install the binary in `$GOBIN` directory.

Now you can run the program with:
```
mainflux-http-sender
```
if `$GOBIN` is in `$PATH` (otherwise use `$GOBIN/mainflux-http-sender`)

### Documentation
Development documentation can be found [here](http://mainflux.io/).

### Community
#### Mailing lists
[mainflux](https://groups.google.com/forum/#!forum/mainflux) Google group.

#### IRC
[Mainflux Gitter](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

#### Twitter
[@mainflux](https://twitter.com/mainflux)

### License
[Apache License, version 2.0](LICENSE)
