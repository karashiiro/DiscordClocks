package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func runClocks(client *discordgo.Session, resources *Resources) {
	for _, clock := range resources.Clocks {
		channel, err := client.Channel(clock.ChannelID)
		if err != nil {
			log.Println(err)
			continue
		}

		channel.Name = "a"
	}
}
