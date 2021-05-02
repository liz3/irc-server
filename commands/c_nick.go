package commands

import "irc-server/models"

func NickCmd() *models.Command {
	return &models.Command{
		Name: "Nick",
		CmdName: "NICK",
		OpCode: "",
		Enabled: true,
		Handler: func(issuer *models.Client, args []string, raw []byte) models.CommandResult {
			if issuer.ConnType == models.ServerConnection {
				return models.InvalidCommand
			}
			if len(args) != 2 || len(args[1]) > 16 {
				return models.InvalidArgument
			}
			var taken = false
			issuer.Instance.RunLocked(func() {
				for _, client := range issuer.Instance.ClientConnections {
					if client.User != nil && client.User.Nick == args[1] {
						taken = true
						break
					}
				}
			})
			if taken {
				issuer.Send(models.ErrNickInUse, []models.Argument{
					models.StarParam(),
					models.SingleParam(args[1], false),
					models.SingleParam("Nickname already in use", true),
			})

			}
			if issuer.User == nil {
				issuer.User = &models.User{
					Nick: args[1],
				}
			} else {
				issuer.User.Nick = args[1]

			}

			return models.Success
		},
	}
}
