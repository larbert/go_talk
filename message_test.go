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
	fmt.Println(data)
	m = BytesToMessage(data)
	fmt.Println(m)
}
