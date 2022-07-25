package main

import (
	"time"

	"github.com/larbert/go_talk"
)

func main() {
	go func() {
		s := &go_talk.Server{
			Name: "111",
		}
		s.Start(":8080")
	}()
	time.Sleep(time.Duration(2) * time.Second)
	c := &go_talk.Client{}
	c.Connect(":8080")
}
