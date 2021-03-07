package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"fmt"
	"os"
	"flag"
	"github.com/jcatala/drgob/pkg/discord"
)



func main() {
	const version = "0.0.1"

	// The items inside flags are just pointers
	// First, parse all the args that are passed to the bot
	verbose := flag.Bool("verbose", false, "To be verbose")
	flag.Parse()

	err := godotenv.Load()
	if err != nil{
		log.Fatalln("Error loading the environment variable " )
	}

	discordkey := ""
	if *verbose{
		discordkey = os.Getenv("discord_key")
		fmt.Printf("Discord key: %s\n", discordkey)
	}


	discord, err := discordgo.New("Bot " + "a")
	if err != nil{
		log.Fatal("Error")
	}

	fmt.Printf("%p", &discord)
	
}
