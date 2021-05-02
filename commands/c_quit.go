package commands

import (
	"irc-server/models"
	"strings"
)

func QuitCmd() *models.Command {
	return &models.Command{
		Name: "Quit",
		CmdName: "QUIT",
		OpCode: "",
		Enabled: true,
		Handler: func(issuer *models.Client, args []string, raw []byte) models.CommandResult {
			var message = ""
			if len(args) > 1 {
				message = strings.Join(args[1:], " ")
			}
			var parts = []string{models.GetUserDescriptor(issuer.User, issuer.Instance.Config), string(models.Quit), ":" + message}
			for _, conn := range issuer.Instance.ClientConnections {
				if conn == issuer {
					continue
				}
				conn.SendRaw(parts)
			}
			issuer.Disconnect()
			issuer.Listener.Close()

			return models.Success
		},
	}
}
