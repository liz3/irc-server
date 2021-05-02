package commands

import (
"irc-server/models"
"strings"
)

func PrivMsgCmd() *models.Command {
	return &models.Command{
		Name: "PrivateMessage",
		CmdName: "PRIVMSG",
		OpCode: "",
		Enabled: true,
		Handler: func(issuer *models.Client, args []string, raw []byte) models.CommandResult {
			if len(args) < 3 {
				return models.InvalidArgument
			}
			var channelName = strings.ToLower(args[1])
			var message = strings.Join(args[2:], " ")
			if message[0] == ':' {
				message = message[1:]
			}

			var channel *models.Channel = nil
			for _, c := range issuer.ActiveChannels {
				if strings.ToLower(c.Name) == channelName {
					channel = c
					break
				}
			}
			if channel == nil {
				return models.InvalidArgument
			}
			var parts = []string{issuer.User.GetUserDescriptor(issuer.Instance.Config), string(models.PrivMsg), channel.Name, ":" + message}

			channel.Broadcast(parts)


			return models.Success
		},
	}
}
