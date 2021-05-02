package commands

import "irc-server/models"

func PingCmd() *models.Command {
	return &models.Command{
		Name: "Ping",
		CmdName: "PING",
		OpCode: "",
		Enabled: true,
		Handler: func(issuer *models.Client, args []string, raw []byte) models.CommandResult {
			issuer.Send(models.Pong, models.SingleParamList(args[1], true))
			return models.Success
		},
	}
}
