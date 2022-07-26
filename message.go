package go_talk

import "unsafe"

type Message struct {
	Option
	Payload string
}

type Option int64

const (
	OptionConnect  Option = 0
	OptionError    Option = 1
	OptionTalk     Option = 2
	OptionACK      Option = 3
	OptionJoinRoom Option = 4
)

// SliceMock 用来转换[]byte和struct
type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

// byte切片转换成Message结构体
func BytesToMessage(data []byte) *Message {
	return *(**Message)(unsafe.Pointer(&data))
}

// Message结构体转换为byte切片
func MessageToBytes(message *Message) []byte {
	messageLen := unsafe.Sizeof(*message)
	return *(*[]byte)(unsafe.Pointer(&SliceMock{
		addr: uintptr(unsafe.Pointer(message)),
		cap:  int(messageLen),
		len:  int(messageLen),
	}))
}
