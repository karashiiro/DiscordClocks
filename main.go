package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	client, err := discordgo.New("Bot " + os.Getenv("DISCLOCKS_BOT_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	resources := Resources{
		Clocks: make([]ClockEntry, 1),
	}
	messageCreate := CreateMessageHandler(&resources)
	client.AddHandler(messageCreate)

	if err = client.Open(); err != nil {
		log.Fatalln(err)
	}

	user, err := client.User("@me")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Logged in as", user.Username)

	go RunClocks(client, &resources)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}
