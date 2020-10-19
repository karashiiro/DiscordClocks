package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// ClockEntry is the basic information needed to run a channel clock.
type ClockEntry struct {
	ChannelID string `json:"channelID"`
	Timezone  string `json:"timezone"`
}

// Resources are shared throughout the application.
type Resources struct {
	Clocks []ClockEntry
}

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

	go runClocks(client, &resources)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}
