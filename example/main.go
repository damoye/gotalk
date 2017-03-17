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
