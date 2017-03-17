# gotalk

Simple Go network communication library.

Gotalk makes up for the missing message boundaries of TCP. It can be combined with serialization tools like JSON and Protocol Buffers. It makes network communication much easier. Its inspiration comes from the [Bulk Strings from RESP](https://redis.io/topics/protocol#resp-bulk-strings)

## Protocol
Gotalk defines a protocol like this:

```go
(length_of_message)\r\n(message)\r\n
```

For example:

message     | bytes
------------|----------------------
hello world | 11\r\nhello world\r\n
hello go    | 8\r\nhello go\r\n

## Installation
```sh
go get github.com/damoye/gotalk
```

## Demo
```go
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/damoye/gotalk"
)

func handleConnection(conn net.Conn) {
	connection := gotalk.NewConnection(conn)
	defer connection.Close()
	for {
		message, err := connection.Receive()
		if err != nil {
			log.Print(err)
			break
		}
		err = connection.Send(message)
		if err != nil {
			log.Print(err)
			break
		}
	}
}

func serve(port int) {
	l, err := net.Listen("tcp", fmt.Sprint(":", port))
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

func main() {
	serve(2000)
}
```
