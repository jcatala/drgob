package config

import (
	"errors"
	"fmt"
	"github.com/jcatala/drgob/pkg/discord"
	"github.com/jcatala/drgob/pkg/reddit"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct{
	Verbose bool	// Config from the own package to be verbose or not
	Prefix string	//	String of the prefix to use
	DiscordThings *discord.DiscordThings	// This will host the discord things, including the bot structure
	RedditThings *reddit.RedditThings		// This will save the reddit things
}

func NewConfig(v bool) (*Config) {
	c := new(Config)
	c.Verbose = v
	return c
}

func (c *Config)ReadOsVariables() (e error){
	err := godotenv.Load()
	if err != nil{
		return errors.New("Error by loading .env variable")
	}
	/*
	Initialize discord bot and reddit things oneshot inside the config variable.
	*/
	c.DiscordThings, err = discord.NewDiscordThings(os.Getenv("discord_key"))
	if err != nil{
		log.Fatal(err.Error())
	}
	if c.Verbose{
		fmt.Printf(`Creating new Discord object with
Token: %s`, c.DiscordThings.DiscordToken)
	}
	id := os.Getenv("reddit_id")
	username := os.Getenv("reddit_username")
	password := os.Getenv("reddit_password")
	secret := os.Getenv("reddit_secret")
	c.RedditThings,err = reddit.NewRedditThings(c.Verbose, id, secret, username, password)
	return err
}
