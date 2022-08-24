package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/misly16/logbot/event"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	  }
	dg, err := discordgo.New(os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatal("error creating Discord session,", err)
		return
	}

	dg.State.MaxMessageCount = 2000
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildMembers)
	event.CreateHandler(dg)

	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}