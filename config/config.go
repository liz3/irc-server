package config

type Config struct {
	Name string
	Indent string
	Version string
	Created string
	Modt string
}

func GetConfig() *Config {
	return &Config {
		Name: "Test Network",
			Indent: "test-server",
			Version: "irc-server 1.0",
			Created: "21:16:27 Mar 25 2021",
			Modt: "this is a cool modt",
	}
}
