package go_talk

import (
	"log"
	"net"
	"unsafe"
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
			log.Printf("Connect success, ip=%v\n", conn, conn.RemoteAddr().String())
			go s.serverProcess(conn)
		}
	}
}

func (s *Server) serverProcess(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024*1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Client out err=", err)
		return
	}
	data := buf[:n]
	message := *(**Message)(unsafe.Pointer(&data))
	log.Println(message)
}
