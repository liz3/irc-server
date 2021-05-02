package models

import (
	"net"
//	"time"
	"fmt"
	"strings"
)

type ConnectionType int
type ClientState int

const (
	ServerConnection ConnectionType = 0
	ClientConnection ConnectionType = 1
)

const (
	Initial ClientState = 0
	ClientReady ClientState = 1
)

type Client struct {
	ConnType ConnectionType
	Listener net.Conn
	User *User
	Server *Server
	State ClientState
	Instance *Instance
	ActiveChannels []*Channel
}
func (c *Client) ExpectRead(size int) ([]byte, bool, error) {
	var expected = make([]byte, size)
	n, err := c.Listener.Read(expected)
	return expected, n == size, err
}

func (c* Client) SendRaw(parts []string) {
	var msgString = strings.Join(parts, " ")
	fmt.Println("Sending", msgString)
	msgString += "\r\n"
	_, err := c.Listener.Write([]byte(msgString))
	if err != nil {
		fmt.Println("err on send", err)
	}

}

func (c *Client) Send(it ICI, args []Argument) {
	var parts []string
	parts = append(parts, ":" + c.Instance.Config.Indent)
	parts = append(parts, string(it))

	for _, arg := range args {
		if arg.IsPrefixed {
			parts = append(parts, ":" + arg.Value)
		} else {
			parts = append(parts, arg.Value)
		}
	}
	c.SendRaw(parts)
}

func (c *Client) setup() {
	// c.Send(Notice, []Argument{
	// 	StarParam(),
	// 	Argument{
	// 		Value: "*** Looking up your ident...",
	// 		IsPrefixed: true,
	// 	},
	// })
}

func (c *Client) Disconnect() {
	c.Instance.RunLocked(func() {
		if c.ConnType == ServerConnection {
			delete(c.Instance.ServerConnections, c.Listener)
		} else {
			delete(c.Instance.ClientConnections, c.Listener)
		}
		for _, curr_channel := range c.ActiveChannels {
			for i, e := range curr_channel.ActiveUsers {
				if e != c {
					continue
				}
				var s = curr_channel.ActiveUsers
				s[len(s)-1], s[i] = s[i], s[len(s)-1]
				curr_channel.ActiveUsers = s[:len(s)-1]
				break
			}
		}

	})

}

func (c *Client) handleMessage(msg [] byte) {
	var raw = string(msg)
	for _, str := range strings.Split(raw, "\r\n") {
		if len(str) == 0 {
			continue
		}
		fmt.Println("Received", str, len(str))
		var parts = strings.Split(str, " ")
		var cmd = c.Instance.Commands.FindByCmd(strings.ToUpper(parts[0]))
		if cmd == nil {
			continue
		}
		_ = cmd.Handler(c, parts, msg)
	}
}

func (c *Client) AbstractHandle() {
	c.setup()
	for {

		var readBytes = make([]byte, 1024)
		n, err := c.Listener.Read(readBytes)
		if err != nil {
			fmt.Println("err on read", err)
			break
		}
		c.handleMessage(readBytes[0:n])
	}
	c.Disconnect()
}
