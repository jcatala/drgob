package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"fmt"
)

func main() {
	const version = "0.0.1"

	err := godotenv.Load()
	if err != nil{
		log.Fatalln("Error loading the environment variable " )
	}


	discord, err := discordgo.New("Bot " + "ODE4MjEwNDg5NTMxMzY3NDM0.YEUv5g.ni1DKEjdvMhDX4QwWDeS_Sr_OSI")
	if err != nil{
		log.Fatal("Error")
	}
	fmt.Printf("%p", &discord)
	
}
