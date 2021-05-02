package models

import (
	"net"
	"sync"
	"irc-server/config"
)

type Instance struct {
	ClientListener *net.Listener
	ServerListener *net.Listener
	ClientConnections map[net.Conn]*Client
	ServerConnections map[net.Conn]*Client
	Channels []*Channel
	Servers []*Server
	Commands *CommandsList
	InstanceMutex *sync.Mutex
	Config *config.Config
	OperatorCount int
}

func (i* Instance) RunLocked(cb func()) {
	i.InstanceMutex.Lock()
	cb()
	i.InstanceMutex.Unlock()
}

func CreateEmptyInstance(clientListener, serverListener *net.Listener) *Instance {
	return &Instance {
		ClientListener: clientListener,
			ServerListener: serverListener,
		ClientConnections: make(map[net.Conn]*Client),
		ServerConnections: make(map[net.Conn]*Client),
		Channels: nil,
		Servers: nil,
		InstanceMutex: &sync.Mutex{},
		Commands: nil,
		OperatorCount: 0,
	}
}
