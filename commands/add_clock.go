package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DiscordClocks/models"
)

// AddClock adds a clock to the clock registry.
func AddClock(client *discordgo.Session, message *discordgo.MessageCreate, args []string, resources *models.Resources) {
	if len(args) < 2 {
		if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<%s>, too few arguments!", message.Author.ID)); err != nil {
			log.Println(err)
		}
		return
	}

	channel := args[0]
	timezone := args[1]
	tzCode := ""
	if len(args) == 3 {
		tzCode = args[2]
	}

	for _, clock := range resources.Clocks {
		if clock.ChannelID == channel {
			if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<%s>, that clock already exists!", message.Author.ID)); err != nil {
				log.Println(err)
			}
			return
		}
	}

	resources.Clocks = append(resources.Clocks, models.ClockEntry{
		ChannelID: channel,
		Timezone:  timezone,
		TzCode:    tzCode,
	})

	resources.Save()

	if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<%s>, the clock was created! Please allow for up to five minutes for the clock to update.", message.Author.ID)); err != nil {
		log.Println(err)
	}
}
