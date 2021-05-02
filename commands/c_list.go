package commands

import (
	"irc-server/models"
	"strings"
	"strconv"
)

func ListCmd() *models.Command {
	return &models.Command{
		Name: "List",
		CmdName: "LIST",
		OpCode: "",
		Enabled: true,
		Handler: func(issuer *models.Client, args []string, raw []byte) models.CommandResult {
			issuer.SendUser(models.RplListStart, []models.Argument{
				models.SingleParam("Channel", false),
				models.SingleParam("Users", true),
				models.SingleParam("False", false),
			})
			var wanted = ""
			if len(args) > 1 {
				wanted = strings.ToLower(args[2])
			}
			for _, channel := range issuer.Instance.Channels {
				if channel.HasFlag(models.CMPrivate) {
					continue
				}
				if len(wanted) > 0 && !strings.Contains(wanted, strings.ToLower(channel.Name)) {
					continue
				}
				issuer.SendUser(models.RplList, []models.Argument{
					models.SingleParam(channel.Name, false),
					models.SingleParam(strconv.Itoa(len(channel.Users)), false),
					models.SingleParam(channel.Topic, true),
				})

			}
			issuer.SendUser(models.RplListEnd, []models.Argument{
				models.SingleParam("End of channel list", true),
			})

			return models.Success
		},
	}
}
