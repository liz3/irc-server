package commands

import "irc-server/models"

func WhoIsCmd() *models.Command {
	return &models.Command{
		Name: "WhoIs",
		CmdName: "WHOIS",
		OpCode: "",
		Enabled: true,
		Handler: func(issuer *models.Client, args []string, raw []byte) models.CommandResult {
			var client *models.Client = nil
			issuer.Instance.RunLocked(func () {
				for _, cc := range issuer.Instance.ClientConnections {
					if cc.User != nil && cc.User.Nick == args[2] {
						client = cc
						break
					}
				}
			})
			if client == nil {
				issuer.Send(models.ErrNoSuchNick, []models.Argument{
					models.SingleParam(issuer.User.Nick, false),
					models.SingleParam(args[1], false),
					models.SingleParam("No such nick/channel", true),
				})

			} else {
				issuer.Send(models.RplWhoIsUser, []models.Argument{
					models.SingleParam(issuer.User.Nick, false),
					models.SingleParam(client.User.Nick, false),
					models.SingleParam(client.User.Username, false),
					models.SingleParam(client.Listener.RemoteAddr().String(), false),
					models.StarParam(),
					models.SingleParam(client.User.Realname, true),
				})
				issuer.Send(models.RplEndWhoIs, []models.Argument{
					models.SingleParam(issuer.User.Nick, false),
					models.SingleParam(client.User.Nick, false),
					models.SingleParam("End of /WHOIS list", true),

				})

			}

			return models.Success
		},
	}
}
