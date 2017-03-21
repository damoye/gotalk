# gotalk
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/damoye/gotalk)
[![Build Status](https://travis-ci.org/damoye/gotalk.svg?branch=master)](https://travis-ci.org/damoye/gotalk)
[![Coverage Status](https://coveralls.io/repos/github/damoye/gotalk/badge.svg)](https://coveralls.io/github/damoye/gotalk)

Simple Go network communication library.

Gotalk makes up for the missing message boundaries of TCP. It can be combined with serialization tools like JSON and Protocol Buffers. It makes network communication much easier. Its inspiration comes from the [Bulk Strings from RESP](https://redis.io/topics/protocol#resp-bulk-strings)

## Protocol
Gotalk defines a protocol like this:

```go
(length_of_message)\r\n(message)\r\n
```

For example:

message       | bytes
--------------|------------------------
"hello world" | "11\r\nhello world\r\n"
""            | "0\r\n\r\n"

## Installation
```sh
go get github.com/damoye/gotalk
```

## Example
```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/damoye/gotalk"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	defer conn.Close()
	for {
		message, err := gotalk.Decode(reader)
		if err != nil {
			log.Print(err)
			break
		}
		_, err = conn.Write(gotalk.Encode(message))
		if err != nil {
			log.Print(err)
			break
		}
	}
}

func main() {
	l, err := net.Listen("tcp", fmt.Sprint(":", 2000))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
```
