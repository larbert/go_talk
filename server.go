package go_talk

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Name string
}

func (s *Server) Start(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("Server start error: ", err)
		return
	}
	defer listener.Close()
	log.Println("Server start success!")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connect error: ", err)
		} else {
			log.Printf("Connect success, ip=%v\n", conn.RemoteAddr().String())
			go s.serverProcess(conn)
		}
	}
}

func (s *Server) serverProcess(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024*1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Client out err=", err)
			return
		}
		data := buf[:n]
		message := BytesToMessage(data)
		log.Println("Serve read success: ", message)
		message = &Message{
			Option: 2,
		}
		_, err = conn.Write(MessageToBytes(message))
		if err != nil {
			log.Println("serve write error")
		} else {
			log.Println("Serve write success")
		}
		fmt.Println("-----------------------------------------------------")
	}
}
