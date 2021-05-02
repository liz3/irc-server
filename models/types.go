package models

import (
	"sync"
	"irc-server/config"
)

type Server struct {

}

type ChannelV string

const (
	PublicChannel ChannelV = "="
	SecretChannel ChannelV = "@"
	PrivateChannel ChannelV = "*"
)

type Channel struct {
	Name string
	Users []*User
	Topic string
	ActiveUsers []*Client
	ChannelLevel ChannelV
}
type User struct {
	Nick string
	Realname string
	UserMode int
	Username string
}

var mutex = &sync.Mutex{}

func (c *Channel) Broadcast(parts []string) {
	for _, client := range c.ActiveUsers {
		client.SendRaw(parts)
	}

}

func (c *Channel) GetUserNames() []string  {
	mutex.Lock()
	var l []string
	for _, u := range c.Users {
		l = append(l, u.Username)
	}

	mutex.Unlock()
	return l
}

func GetUserDescriptor(user *User, config *config.Config) string {
	return user.Nick + "!" + user.Username + "@" +config.Indent
}
