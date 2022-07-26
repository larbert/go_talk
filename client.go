package go_talk

import (
	"log"
	"net"
)

type Client struct {
}

func (c *Client) Connect(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln("Connect err: ", err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024*1024)
	for {
		message := &Message{
			Option: 0,
		}
		_, err := conn.Write(MessageToBytes(message))
		if err != nil {
			log.Println("Client.write error: ", err)
		} else {
			log.Println("Client write success")
		}
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Client read error: ", err)
		} else {
			data := buf[:n]
			message = BytesToMessage(data)
			log.Println("Client read success: ", message)
		}
	}
}
