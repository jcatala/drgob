package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"fmt"
	"os"
	"flag"
	"github.com/jcatala/drgob/pkg/config"
	"github.com/jcatala/drgob/pkg/discord"
	"errors"
	"os/signal"
	"syscall"
)


func getDcToken() (response string, e error){
	err := godotenv.Load()
	if err != nil{
		return "", errors.New("Error by loading .env variable")
	}
	dcToken := os.Getenv("discord_key")
	return dcToken, nil
}


func main() {
	const version = "0.0.1"

	// The items inside flags are just pointers
	// First, parse all the args that are passed to the bot
	verbose := flag.Bool("verbose", false, "To be verbose")
	flag.Parse()

	// Creating a config options to push everything inside
	//var config *config.Config
	config := config.NewConfig(*verbose)
	// Loading environment variables, including reddit and dc tokens

	// First dc token
	dcToken, err := getDcToken()
	if err != nil{
		log.Fatalln("Error getting .env")
	}
	if *verbose{
		fmt.Printf("Discord key: %s\n", dcToken)
	}
	// Create a discordThing object with the dc token
	config.DiscordThings, err = discord.NewDiscordThings(dcToken)
	if err != nil{
		log.Fatalln("Error on initialization of the discord bot")
	}

	// Now, we have the config created with the discord object correctly created, we now need to create a handler
	config.DiscordThings.DiscordSession.AddHandler(config.DiscordThings.TestMessage)
	config.DiscordThings.DiscordSession.Identify.Intents = discordgo.IntentsGuildMessages
	err = config.DiscordThings.DiscordSession.Open()
	if err != nil{
		log.Fatalln("Error on oppening discord thing")
	}
	// Wait here until CTRL + C
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// If signal, close everything flawlessly
	config.DiscordThings.DiscordSession.Close()

}
