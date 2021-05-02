package commands

import (
"irc-server/models"
"strings"
)

func TopicCmd() *models.Command {
	return &models.Command{
		Name: "Channel Topic command",
		CmdName: "TOPIC",
		OpCode: "",
		Enabled: true,
		Handler: func(issuer *models.Client, args []string, raw []byte) models.CommandResult {
			if len(args) < 2 {
				return models.InvalidArgument
			}
			var channelName = strings.ToLower(args[1])
			var message = ""
			var set = len(args) > 2
			if set {
				if args[2] == ":" {
					message = ":"
				} else {
					message = strings.Join(args[2:], " ")
					if message[0] == ':' {
						message = message[1:]
					}

				}
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

			if !set {
				if len(channel.Topic) == 0 {
				issuer.Send(models.RplNoTopic, []models.Argument{
					models.SingleParam(channel.Name, false),
					models.SingleParam("No topic set", true),
				})

				} else {
				issuer.Send(models.RplTopic, []models.Argument{
					models.SingleParam(channel.Name, false),
					models.SingleParam(channel.Topic, true),
				})
				}
			} else {
				if message == ":" {
					channel.Topic = ""
				} else {
					channel.Topic = message
				}
			}


			return models.Success
		},
	}
}
