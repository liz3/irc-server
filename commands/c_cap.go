package commands

import "irc-server/models"

func CapCmd() *models.Command {
	return &models.Command{
		Name: "CapCmd",
		CmdName: "CAP",
		OpCode: "",
		Enabled: true,
		Handler: func(issuer *models.Client, args []string, raw []byte) models.CommandResult {
			// TODO!
			return models.Success
		},
	}
}
