package go_talk

import (
	"fmt"
	"testing"
)

func TestMessage(t *testing.T) {
	m := &Message{
		Option: 1,
	}
	data := MessageToBytes(m)
	fmt.Println(data)
	m = BytesToMessage(data)
	fmt.Println(m)
}
