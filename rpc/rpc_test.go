package rpc

import (
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("expected: %s, actual: %s", expected, actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	incoming := "Content-Length: 23\r\n\r\n{\"Method\":\"somemethod\"}"
	method, content, err := DecodeMessage([]byte(incoming))
	contentLength := len(content)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if contentLength != 23 {
		t.Fatalf("expected: 23, got: %d", contentLength)
	}

	if method != "somemethod" {
		t.Fatalf("expected: somemethod, got: %s", method)
	}
}
