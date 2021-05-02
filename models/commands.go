package models


type CommandResult int

const (
	Success CommandResult = 0
	InvalidCommand CommandResult = 1
	InvalidArgument CommandResult = 2
)

type Command struct {
	Name string
	CmdName string
	OpCode string
	Handler func(*Client, []string,  []byte) CommandResult
	Enabled bool
}
type CommandsList struct {
	Commands []*Command

}
func (l *CommandsList) FindByCmd(name string) *Command {
	for _, entry := range l.Commands {
		if entry.CmdName == name {
			return entry
		}
	}
	return nil
}
func (l *CommandsList) FindByOp(name string) *Command {
	for _, entry := range l.Commands {
		if entry.OpCode == name {
			return entry
		}
	}
	return nil
}
