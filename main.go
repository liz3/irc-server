package main

import (
	"irc-server/models"
	"fmt"
	"irc-server/server"
	"irc-server/commands"
	"irc-server/config"
	"net"
	"time"
)

func main() {
	fmt.Println("Starting listeners")
	clientListener := server.CreateServer(":8080")
	serverListener := server.CreateServer(":8081")
	var instance = models.CreateEmptyInstance(clientListener, serverListener)
	instance.Commands = commands.GetCommands()
	instance.Config = config.GetConfig()
	fmt.Println("starting to listen for connections")
	go server.HandleConnections(*clientListener, func(c net.Conn) {
		var client = server.CreateClient(c, instance, models.ClientConnection)
		instance.RunLocked(func() {
			instance.ClientConnections[c] = client
		})
	})
	go server.HandleConnections(*serverListener, func(c net.Conn) {
		var client = server.CreateClient(c, instance, models.ServerConnection)
		instance.RunLocked(func() {
			instance.ServerConnections[c] = client
		})



	})
	for {
		time.Sleep(200 * time.Millisecond)

	}
}
