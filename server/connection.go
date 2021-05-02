package server

import (
	"net"
	"irc-server/models"
)

func CreateClient(listener net.Conn, instance *models.Instance, ctype models.ConnectionType) *models.Client {
	var client = &models.Client{
		ConnType: ctype,
		Listener: listener,
		User: nil,
		Server: nil,
		State: models.Initial,
		Instance: instance,
	}
	go client.AbstractHandle()
	return client
}
