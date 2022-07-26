package main

import (
	"fmt"

	"github.com/larbert/go_talk"
)

func main() {
	menu()
	var op int
	fmt.Scanf("%d", &op)
	switch op {
	case 1:
		serve(":8080", "1111")
	case 2:
		var addr string
		fmt.Print("Server IP: ")
		fmt.Scanf("%s", &addr)
		client(addr)
	}
}

func serve(addr string, name string) {
	s := &go_talk.Server{
		Name: "name",
	}
	s.Start(addr)
}

func client(addr string) {
	c := &go_talk.Client{}
	c.Connect(addr)
}

func menu() {
	fmt.Println("**************************")
	fmt.Println("* 1. server              *")
	fmt.Println("* 2. client              *")
	fmt.Println("**************************")
}
