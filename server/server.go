package server

import (
	"net"
	"fmt"
)


func CreateServer(port string) *net.Listener {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil
	}
	return &ln
}
func HandleConnections(listener net.Listener, callback func(net.Conn)) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err", err)
			continue
		}
		fmt.Println("New connection")
		callback(conn)
	}
}
