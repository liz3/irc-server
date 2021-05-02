package commands

import (
	"irc-server/models"
	"irc-server/util"
	"strings"
	"strconv"
)

func sendModt(client *models.Client) {
		var config = client.Instance.Config
	client.SendUser(models.RplModtStart, []models.Argument{
		models.SingleParam(config.Indent + " Message of the day", true),
	})
	for _, e := range strings.Split(config.Modt, "\n") {
		for _,p := range util.SplitMessageIrc(e) {
			client.SendUser(models.RplModt, []models.Argument{
				models.SingleParam(p, true),
			})
		}
	}
	client.SendUser(models.RplModtEnd, []models.Argument{
		models.SingleParam("End Message of the day", true),
	})
}

func sendWelcome(client *models.Client) {
	var config = client.Instance.Config
	client.SendUser(models.RplWelcome, []models.Argument{
		models.SingleParam("Welcome to the 42Net IRC Network " + client.User.Nick + "!" + client.User.Username + "@" + config.Indent, true),
	})

	client.SendUser(models.RplYourHost, []models.Argument{
		models.SingleParam("Your host is " + config.Indent + ", running version " + config.Version, true),
	})
	client.SendUser(models.RplCreated, []models.Argument{
		models.SingleParam("Your server was created: " + config.Created, true),
	})

	client.SendUser(models.RplMyInfo, []models.Argument{
		models.SingleParam(config.Name + " " + config.Version + " 0123456789 HILRSchiorsw ACDFHIJKLMOPQRSYabcdeghijklmnopqrstuvz :FHIJLYabdeghjkloqv", false),
	})

	client.SendUser(models.RplLUserClient, []models.Argument{
		models.SingleParam("There are " + strconv.Itoa(len(client.Instance.ClientConnections)) + " on 1 server", true),
	})
	client.SendUser(models.RplLUserOp, []models.Argument{
		models.SingleParam("There are " + strconv.Itoa(client.Instance.OperatorCount) + " :operators online", false),
	})

	client.SendUser(models.RplLChannels, []models.Argument{
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
