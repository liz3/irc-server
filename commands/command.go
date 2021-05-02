package commands

import (
	"irc-server/models"
)
func GetCommands() *models.CommandsList {
	var cmdList []*models.Command

	cmdList = append(cmdList, NickCmd())
	cmdList = append(cmdList, UserCmd())
	cmdList = append(cmdList, PingCmd())
	cmdList = append(cmdList, JoinCmd())
	cmdList = append(cmdList, PrivMsgCmd())
	cmdList = append(cmdList, TopicCmd())
	cmdList = append(cmdList, ListCmd())
	cmdList = append(cmdList, WhoIsCmd())
	cmdList = append(cmdList, QuitCmd())
	return &models.CommandsList {
		Commands: cmdList,
	}
}
