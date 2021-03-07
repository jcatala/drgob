package config

import "github.com/jcatala/drgob/pkg/discord"

type Config struct{
	Verbose bool	// Config from the own package to be verbose or not
	DiscordThings *discord.DiscordThings	// This will host the discord things, including the bot structure
}

func NewConfig(v bool) (*Config) {
	c := new(Config)
	c.Verbose = v
	return c
}
