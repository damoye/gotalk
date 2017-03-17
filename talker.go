package gotalk

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
)

// Connection is a message wrapper for net.Conn
type Connection struct {
	conn   net.Conn
	reader *bufio.Reader
}

// NewConnection returns a new Connection
func NewConnection(conn net.Conn) Connection {
	return Connection{conn, bufio.NewReader(conn)}
}

// Send send a message. It can be called by multiple goroutine.
func (connection *Connection) Send(message string) error {
	_, err := fmt.Fprintf(connection.conn, "%d\r\n%s\r\n", len(message), message)
	if err != nil {
		return err
	}
	return nil
}

// Receive receive a message. It can only be called by ONE goroutine.
func (connection *Connection) Receive() (message string, err error) {
	head, err := connection.reader.ReadString('\n')
	if err != nil {
		return
	}
	length, err := strconv.Atoi(head[:len(head)-2])
	if err != nil {
		return
	}
	buff := make([]byte, length+2)
	_, err = io.ReadFull(connection.reader, buff)
	if err != nil {
		return
	}
	return string(buff[:length]), nil
}

// Close close the underlying net.Conn
func (connection *Connection) Close() {
	connection.conn.Close()
}
