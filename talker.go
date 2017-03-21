package gotalk

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Encode encodes the message
func Encode(message string) []byte {
	return []byte(fmt.Sprintf("%d\r\n%s\r\n", len(message), message))
}

// Decode decode from reader
func Decode(reader *bufio.Reader) (message string, err error) {
	head, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	if len(head) < 3 {
		err = fmt.Errorf("head %q is too short", head)
		return
	}
	length, err := strconv.Atoi(head[:len(head)-2])
	if err != nil {
		return
	}
	buff := make([]byte, length+2)
	_, err = io.ReadFull(reader, buff)
	if err != nil {
		return
	}
	return string(buff[:length]), nil
}
