package go_talk

import "unsafe"

type Message struct {
	Option
	Payload interface{}
}

type Option int64

const (
	Connect Option = 0
	Talk    Option = 1
)

func BytesToMessage(data []byte) *Message {
	return *(**Message)(unsafe.Pointer(&data))
}

func MessageToBytes(message *Message) []byte {
	return *(*[]byte)(unsafe.Pointer(message))
}
