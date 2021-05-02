package models

type Argument struct {
	Value string
	IsPrefixed bool

}

type ICI string

const (
	Notice ICI = "NOTICE"
	Pong ICI = "PONG"
	Join ICI = "JOIN"
	PrivMsg ICI = "PRIVMSG"
	Quit ICI = "QUIT"
	RplWelcome ICI = "001"
	RplYourHost ICI = "002"
	RplCreated ICI = "003"
	RplMyInfo ICI = "004"
	RplLUserClient ICI = "251"
	RplLUserOp ICI = "252"
	RplLChannels ICI = "254"
	RplModtStart ICI = "375"
	RplWhoIsUser ICI = "311"
	RplEndWhoIs ICI = "318"
	RplModt ICI = "372"
	RplModtEnd ICI = "376"
	RplNoTopic ICI = "331"
	RplTopic ICI = "332"
	RplNameReply ICI = "353"
	RplEndNames ICI = "366"
	RplListStart ICI = "321"
	RplList ICI = "322"
	RplListEnd ICI = "323"
	ErrNickInUse ICI = "433"
	ErrNoSuchNick ICI = "401"
)

type UserMode string

const (
	MAway UserMode = "a"
	MInvisible UserMode = "i"
	MWallOps UserMode = "w"
	MRestricted UserMode = "r"
	MOperator UserMode = "o"
	MLocalOperator UserMode = "O"
	MReceiveNotices UserMode = "s"
)

type ChannelMode string
type ChannelUserMode string
const (
	// channel flags
	CMAnonymous ChannelMode = "a"
	CMInviteOnly ChannelMode = "i"
	CMModerated ChannelMode = "m"
	CMNoMessages ChannelMode = "n"
	CMQuiet ChannelMode = "q"
	CMPrivate ChannelMode = "p"
	CMSecret ChannelMode = "s"
)

const (
	CMUCreator ChannelUserMode = "q"
	CMUOperator ChannelUserMode = "o"
	CMUVoice ChannelUserMode = "v"
	CMUModerator ChannelUserMode = "h"
)

func GetCMUPrefix(mode ChannelUserMode) string {
	switch mode {
	case CMUCreator:
		return "~"
	case CMUOperator:
		return "@"
	case CMUModerator:
		return "%"
	case CMUVoice:
		return "+"
	default:
		return ""
	}
}

func StarParam() Argument {
	return Argument{
		Value: "*",
		IsPrefixed: false,
	}
}
func SingleParamList(val string, prefixed bool) []Argument {
	return []Argument{
		Argument{
			Value: val,
			IsPrefixed: prefixed,
		},
	}
}
func SingleParam(val string, prefixed bool) Argument {
	return Argument{
			Value: val,
			IsPrefixed: prefixed,
		}
}
