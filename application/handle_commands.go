package application

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DiscordClocks/commands"
)

// CreateMessageHandler curries the message creation delegate with the provided application resources.
func CreateMessageHandler(resources *models.Resources) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(client *discordgo.Session, message *discordgo.MessageCreate) {
		messageCreateInternal(client, message, resources)
	}
}

func messageCreateInternal(client *discordgo.Session, message *discordgo.MessageCreate, resources *models.Resources) {
	if message.Author.Bot {
		return
	}

	member, err := client.GuildMember(message.GuildID, message.Author.ID)
	if err != nil {
		log.Println(err)
		return
	}

	roleOk := false
	for _, role := range member.Roles {
		for _, modRole := range resources.ModRoles {
			if modRole == role {
				roleOk = true
				break
			}
		}
	}
	if !roleOk {
		return
	}

	if message.Content[0:1] != resources.Prefix {
		return
	}

	content := message.Content[1:]

	if strings.HasPrefix(content, "addclock") {
		commands.AddClock(client, message, resources)
	} else if strings.HasPrefix(content, "removeclock") {
		commands.RemoveClock(client, message, resources)
	}
}
