package commands

import (
	"irc-server/models"
	"strings"
	"irc-server/util"
)

func JoinCmd() *models.Command {
	return &models.Command{
		Name: "Join Command",
		CmdName: "JOIN",
		OpCode: "",
		Enabled: true,
		Handler: func(issuer *models.Client, args []string, raw []byte) models.CommandResult {
			if len(args) < 2 {
				return models.InvalidArgument
			}
			var channelName = strings.ToLower(args[1])
			var channel *models.Channel = nil
			issuer.Instance.RunLocked(func() {
				for _, c := range issuer.Instance.Channels {
					if c.Name == channelName {
						channel = c
						break
					}
				}
				if channel == nil {
					channel = &models.Channel{
						Name: channelName,
						Users: nil,
						ActiveUsers: nil,
						Topic: "",
						ChannelLevel: models.SecretChannel,
					}
					issuer.Instance.Channels = append(issuer.Instance.Channels, channel)
				}
				channel.Users = append(channel.Users, issuer.User)
				channel.ActiveUsers = append(channel.ActiveUsers, issuer)
				issuer.ActiveChannels = append(issuer.ActiveChannels, channel)
			})

			channel.Broadcast([]string{":" + models.GetUserDescriptor(issuer.User, issuer.Instance.Config), string(models.Join), ":" + channelName})

			issuer.Send(models.Join, []models.Argument{
				models.SingleParam(channelName, true),
			})
			if len(channel.Topic) == 0 {
				issuer.Send(models.RplNoTopic, []models.Argument{
					models.SingleParam(channelName, false),
					models.SingleParam("No Channel Topic Set", true),
				})

			} else {
				issuer.Send(models.RplTopic, []models.Argument{
					models.SingleParam(channelName, false),
					models.SingleParam(channel.Topic, true),
				})
			}
			var userNames = strings.Join(channel.GetUserNames(), " ")

			for _, entry := range util.SplitMessageIrc(userNames) {
				issuer.Send(models.RplNameReply, []models.Argument{
					models.SingleParam(issuer.User.Username, false),
					models.SingleParam(string(channel.ChannelLevel), false),
					models.SingleParam(channel.Name, false),
					models.SingleParam(entry, true),
				})
			}
			issuer.Send(models.RplEndNames, []models.Argument{
				models.SingleParam(issuer.User.Username, false),
				models.SingleParam(channel.Name, false),
				models.SingleParam("End of /NAMES List", true),
			})


			return models.Success
		},
	}
}
