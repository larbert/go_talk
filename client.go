package go_talk

import (
	"fmt"
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
		var payload string
		var option int64
		fmt.Println("----Options----")
		fmt.Println("|OptionConnect |")
		fmt.Println("----Options----")
		fmt.Println("option")
		fmt.Scanf("%d", &option)
		fmt.Println("payload")
		fmt.Scanf("%s", &payload)
		message := &Message{
			Option:  Option(option),
			Payload: payload,
		}
		log.Println(message)
		log.Println(MessageToBytes(message))
		log.Println(BytesToMessage(MessageToBytes(message)))
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
			switch message.Option {
			case OptionACK:
				log.Println("ACK")
			}
		}
	}
}
