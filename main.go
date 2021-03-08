package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jcatala/drgob/pkg/config"
	"github.com/joho/godotenv"
	"github.com/vartanbeno/go-reddit/reddit"
	"log"
	"os"
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
	prefix := flag.String("prefix", ";", "Prefix to use, default: ;")
	flag.Parse()

	// Creating a config options to push everything inside
	//var config *config.Config
	config := config.NewConfig(*verbose)
	// Loading environment variables, including reddit and dc tokens
	err := config.ReadOsVariables()
	if err != nil{
		log.Fatal(err.Error())
	}

	// add the prefix to the config
	config.Prefix = *prefix
	// Now, we have the config created with the discord object correctly created, we now need to create a handler
	//config.DiscordThings.DiscordSession.AddHandler(config.DiscordThings.TestMessage)
	config.DiscordThings.DiscordSession.AddHandler(config.GetRandomPost)
	config.DiscordThings.DiscordSession.Identify.Intents = discordgo.IntentsGuildMessages
	err = config.DiscordThings.DiscordSession.Open()
	if err != nil{
		log.Fatalln("Error on oppening discord thing")
	}
	test(config)
	// Wait here until CTRL + C
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// If signal, close everything flawlessly
	config.DiscordThings.DiscordSession.Close()

}

func test(c *config.Config){
	posts, _, err := c.RedditThings.RedditClient.Subreddit.TopPosts(context.Background(),
		"battlestations",
		&reddit.ListPostOptions{ListOptions: reddit.ListOptions{Limit: 5},Time: "all"})
	if err != nil{
		log.Fatalln(err.Error())
	}
	fmt.Println("Received posts:")
	fmt.Println(posts[0].Permalink)
}
