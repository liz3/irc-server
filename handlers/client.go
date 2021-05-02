package handlers

import (
	"irc-server/models"
	"fmt"

)

func (c *Client) AbstractHandle() {
	fmt.Println("test")
	for {
		time.Sleep(50 * time.Millisecond)

	}

}
