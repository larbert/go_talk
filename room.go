package go_talk

import "net"

type Room struct {
	Name  string
	Conns []net.Conn
}

func (r *Room) Join(conn net.Conn) {
	r.Conns = append(r.Conns, conn)
}
