package commands

import (
	"irc-server/models"
	"irc-server/util"
	"strings"
	"strconv"
)

func sendModt(client *models.Client) {
		var config = client.Instance.Config
	client.Send(models.RplModtStart, []models.Argument{
		models.SingleParam(client.User.Nick, false),
		models.SingleParam(config.Indent + " Message of the day", true),
	})
	for _, e := range strings.Split(config.Modt, "\n") {
		for _,p := range util.SplitMessageIrc(e) {
			client.Send(models.RplModt, []models.Argument{
				models.SingleParam(client.User.Nick, false),
				models.SingleParam(p, true),
			})
		}
	}
	client.Send(models.RplModtEnd, []models.Argument{
		models.SingleParam(client.User.Nick, false),
		models.SingleParam("End Message of the day", true),
	})
}

func sendWelcome(client *models.Client) {
	var config = client.Instance.Config
	client.Send(models.RplWelcome, []models.Argument{
		models.SingleParam(client.User.Nick, false),
		models.SingleParam("Welcome to the 42Net IRC Network " + client.User.Nick + "!" + client.User.Username + "@" + config.Indent, true),
	})

	client.Send(models.RplYourHost, []models.Argument{
		models.SingleParam(client.User.Nick, false),
		models.SingleParam("Your host is " + config.Indent + ", running version " + config.Version, true),
	})
	client.Send(models.RplCreated, []models.Argument{
		models.SingleParam(client.User.Nick, false),
		models.SingleParam("Your server was created: " + config.Created, true),
	})

	client.Send(models.RplMyInfo, []models.Argument{
		models.SingleParam(client.User.Nick, false),
		models.SingleParam(config.Name + " " + config.Version + " 0123456789 HILRSchiorsw ACDFHIJKLMOPQRSYabcdeghijklmnopqrstuvz :FHIJLYabdeghjkloqv", false),
	})

	client.Send(models.RplLUserClient, []models.Argument{
		models.SingleParam(client.User.Nick, false),
		models.SingleParam("There are " + strconv.Itoa(len(client.Instance.ClientConnections)) + " on 1 server", true),
	})
	client.Send(models.RplLUserOp, []models.Argument{
		models.SingleParam(client.User.Nick, false),
		models.SingleParam("There are " + strconv.Itoa(client.Instance.OperatorCount) + " :operators online", false),
	})

	client.Send(models.RplLChannels, []models.Argument{
		models.SingleParam(client.User.Nick, false),
		models.SingleParam(strconv.Itoa(len(client.Instance.Channels)) + " :channels formed", false),
	})


}

func UserCmd() *models.Command {
	return &models.Command{
		Name: "UserCommand",
		CmdName: "USER",
		OpCode: "",
		Enabled: true,
		Handler: func(issuer *models.Client, args []string, raw []byte) models.CommandResult {
			if issuer.ConnType == models.ServerConnection {
				return models.InvalidCommand
			}
			if len(args) < 4 {
				return models.InvalidArgument
			}
			var username string
			var usermode int
			if args[1] == "*" {
				username = ""
			} else {
				username = args[1]
			}
			if args[2] == "*" {
				usermode = 0
			}else {
				convert, err := strconv.Atoi(args[2])
				if err != nil {
					return models.InvalidArgument
				}
				usermode = convert
			}
			realname := strings.Join(args[4:], " ")
			if issuer.User == nil {
				issuer.User = &models.User{
					Nick: "",
					Realname: realname,
					UserMode: usermode,
					Username: username,
				}
			} else {
				issuer.User.Realname = realname
				issuer.User.UserMode = usermode
				issuer.User.Username = username
			}
			issuer.State = models.ClientReady
			sendWelcome(issuer)
			sendModt(issuer)
			return models.Success
		},
	}
}
