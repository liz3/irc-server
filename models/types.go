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

type ChannelModeEntry struct {
	Flag ChannelMode
	M byte
}
type ChannelUserModeEntry struct {
	Flag ChannelUserMode
	Nick string
	M byte
}

type Channel struct {
	Name string
	Users []*User
	Topic string
	ActiveUsers []*Client
	ChannelModes []ChannelModeEntry
	ChannelUserModes []ChannelUserModeEntry
}

func (channel *Channel) UserHasFlag(user *User, mode ChannelUserMode) bool {
	if mode == CMUCreator {
		for _, e := range channel.ChannelUserModes {
			if e.Flag == CMUCreator && e.Nick == user.Nick && e.M == '+' {
				return true
			}
		}
	} else {
		for _, e := range channel.ChannelUserModes {
			if e.Flag == mode && e.Nick == user.Nick && e.M == '+' {
				return true
			}
		}

	}
	return false
}

func (channel *Channel) HasFlag(mode ChannelMode) bool {
	for _, e := range channel.ChannelModes {
		if e.Flag == mode && e.M == '+' {
			return true
		}
	}
	return false
}

func (channel *Channel) GetChannelMode() string {
	if(channel.HasFlag(CMPrivate)) {
		return "*"
	}
	if(channel.HasFlag(CMSecret)) {
		return "@"
	}
	return "="
}

type UserModeEntry struct {
	Flag UserMode
	M byte
}

type User struct {
	Nick string
	Realname string
	UserMode int
	Username string
	Modes []UserModeEntry
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
		if c.UserHasFlag(u,CMUOperator) {
			l = append(l, "@" + u.Username)
		} else {
			l = append(l, u.Username)
		}

	}

	mutex.Unlock()
	return l
}

func (user *User) GetUserDescriptor(config *config.Config) string {
	return user.Nick + "!" + user.Username + "@" +config.Indent

}
func (user *User) HasFlag(flag UserMode) bool {
	for _, uF := range user.Modes {
		if uF.Flag == flag && uF.M == '+' {
			return true
		}
	}
	return false
}
func (user *User) AddMode(mode UserMode) {
	for index, uF := range user.Modes {
		if uF.Flag == mode{
			user.Modes[index] = UserModeEntry {
				Flag: mode,
					M: '+',
				}
		}
	}
	user.Modes = append(user.Modes, UserModeEntry {
		Flag: mode,
			M: '+',
		},
	)

}
