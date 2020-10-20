package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DiscordClocks/application"
	"github.com/karashiiro/DiscordClocks/models"
)

func main() {
	client, err := discordgo.New("Bot " + os.Getenv("DISCLOCKS_BOT_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	resources := models.LoadResources()
	messageCreate := application.CreateMessageHandler(resources)
	client.AddHandler(messageCreate)

	if err = client.Open(); err != nil {
		log.Fatalln(err)
	}

	user, err := client.User("@me")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Logged in as", user.Username+"#"+user.Discriminator)

	go application.RunClocks(client, resources)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}
