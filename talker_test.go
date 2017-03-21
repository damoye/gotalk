package gotalk

import (
	"bufio"
	"strings"
	"testing"
)

var cases = [][2]string{
	{"hello world", "11\r\nhello world\r\n"},
	{"", "0\r\n\r\n"},
}

func TestEncode(t *testing.T) {
	for _, item := range cases {
		if result := string(Encode(item[0])); result != item[1] {
			t.Errorf("Encode %q return %q, want %q", item[0], result, item[1])
		}
	}
}

func TestDecode(t *testing.T) {
	for _, item := range cases {
		result, err := Decode(bufio.NewReader(strings.NewReader(item[1])))
		if err != nil || result != item[0] {
			t.Errorf("Decode %q return %q, %v, want %q, %v", item[1], result, err, item[0], nil)
		}
	}
	multipleCase := ""
	for _, item := range cases {
		multipleCase += item[1]
	}
	reader := bufio.NewReader(strings.NewReader(multipleCase))
	for _, item := range cases {
		result, err := Decode(reader)
		if err != nil || result != item[0] {
			t.Errorf("Decode %q return %q, %v, want %q, %v", item[1], result, err, item[0], nil)
		}
	}
}

func TestDecodeError(t *testing.T) {
	badCases := []string{
		"",
		"nocrlf",
		"\r\nshorthead",
		"badint\r\nmessage",
		"10\r\nshortbody\r\n",
	}
	for _, item := range badCases {
		_, err := Decode(bufio.NewReader(strings.NewReader(item)))
		if err == nil {
			t.Errorf("Decode %q should failed", item)
		}
	}
}
