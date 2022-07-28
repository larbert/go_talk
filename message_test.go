package go_talk

import (
	"fmt"
	"testing"
)

func TestMessage(t *testing.T) {
	m := &Message{
		Option:  5,
		Payload: "Lyg",
	}
	fmt.Println(m)
	data := MessageToBytes(m)
	data = [5 0 0 0 0 0 0 0 32 161 1 0 192 0 0 0 3 0 0 0 0 0 0 0]
	fmt.Println(data)
	m = BytesToMessage(data)
	fmt.Println(m)
}
