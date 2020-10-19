package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DiscordClocks/models"
)

// RunClocks runs the clocks.
func RunClocks(client *discordgo.Session, resources *models.Resources) {
	for _, clock := range resources.Clocks {
		channel, err := client.Channel(clock.ChannelID)
		if err != nil {
			log.Println(err)
			continue
		}

		channel.Name = "a"
	}
}
