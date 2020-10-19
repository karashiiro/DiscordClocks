package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DiscordClocks/application"
)

func runClocks(client *discordgo.Session, resources *application.Resources) {
	for _, clock := range resources.Clocks {
		channel, err := client.Channel(clock.ChannelID)
		if err != nil {
			log.Println(err)
			continue
		}

		channel.Name = "a"
	}
}
