package go_talk

import (
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
	log.Println("Server start success at ", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connect error: ", err)
		} else {
			log.Printf("Connect success, ip=%v\n", conn.RemoteAddr().String())
			go serverProcess(conn)
		}
	}
}

func serverProcess(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024*1024)
	for {
		// 读取并解析Message
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Client out err=", err)
			return
		}
		data := buf[:n]
		message := BytesToMessage(data)
		log.Println("Serve read success: ", message)
		// 根据Message的Option执行对应的操作
		response := &Message{}
		switch message.Option {
		case OptionConnect:
			connect(conn, response)
		case OptionTalk:
		case OptionJoinRoom:
		default:
			response.Option = OptionError
			response.Payload = "Option error"
		}
		// 返回响应
		_, err = conn.Write(MessageToBytes(response))
		if err != nil {
			log.Println("serve write error")
		} else {
			log.Println("Serve write success")
		}
	}
}

func connect(conn net.Conn, response *Message) {
	response.Option = OptionACK
}
