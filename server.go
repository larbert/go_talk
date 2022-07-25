package go_talk

import (
	"net"
)

type Server struct {
	Name string
	Ip   string
}

func (s *Server) Start(port string) {
	listener := net.Listen("tcp", port)

}
