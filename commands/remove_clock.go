package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DiscordClocks/models"
)

// RemoveClock removes a clock from the clock registry.
func RemoveClock(client *discordgo.Session, message *discordgo.MessageCreate, args []string, resources *models.Resources) {
	if len(args) < 1 {
		if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<@%s>, too few arguments!", message.Author.ID)); err != nil {
			log.Println(err)
		}
		return
	}

	channel := args[0]

	i := -1
	for j, clock := range resources.Clocks {
		if clock.ChannelID == channel {
			i = j
		}
	}
	if i != -1 {
		resources.Clocks = *splice(resources.Clocks, i)
	} else {
		if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<@%s>, no clock exists for that channel!", message.Author.ID)); err != nil {
			log.Println(err)
		}
		return
	}

	resources.Save()

	log.Println("Clock removed from channel", message.ChannelID)
	if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<@%s>, the clock was removed!", message.Author.ID)); err != nil {
		log.Println(err)
	}
}

func splice(slice []models.ClockEntry, i int) *[]models.ClockEntry {
	newSlice := make([]models.ClockEntry, len(slice)-1)
	copy(newSlice, slice[0:i])
	copy(newSlice[i:], slice[i+1:])
	return &newSlice
}
