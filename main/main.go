package main

import (
	"fmt"
	"os"

	"github.com/larbert/go_talk"
)

func main() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
	if len(os.Args) == 4 {
		if os.Args[1] == "-s" {
			serve(os.Args[2], os.Args[3])
		}
	}
	menu()
	var op int
	fmt.Scanf("%d", &op)
	switch op {
	case 1:
		var port string
		var name string
		fmt.Print("Server addr: ")
		fmt.Scanf("%s", &port)
		fmt.Print("Server name: ")
		fmt.Scanf("%s", &name)
		serve(port, name)
	case 2:
		var addr string
		fmt.Print("Server IP: ")
		fmt.Scanf("%s", &addr)
		fmt.Println("Read to connect ", addr)
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
