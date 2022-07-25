package go_talk

import "unsafe"

type Message struct {
	Option
	Payload interface{}
}

type Option int64

const (
//Connect Option = 0
//Talk    Option = 1
)

// SliceMock 用来转换[]byte和struct
type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func BytesToMessage(data []byte) *Message {
	return *(**Message)(unsafe.Pointer(&data))
}

func MessageToBytes(message *Message) []byte {
	messageLen := unsafe.Sizeof(*message)
	return *(*[]byte)(unsafe.Pointer(&SliceMock{
		addr: uintptr(unsafe.Pointer(message)),
		cap:  int(messageLen),
		len:  int(messageLen),
	}))
}
