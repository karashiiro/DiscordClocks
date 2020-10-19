package application

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DiscordClocks/commands"
	"github.com/karashiiro/DiscordClocks/models"
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

	if !rolesOk(member.Roles, resources.ModRoles) {
		return
	}

	if message.Content[0:1] != resources.Prefix {
		return
	}

	content := message.Content[1:]
	args := strings.Split(content, " ")

	if strings.HasPrefix(content, "addclock") {
		commands.AddClock(client, message, args, resources)
	} else if strings.HasPrefix(content, "removeclock") {
		commands.RemoveClock(client, message, args, resources)
	}
}

func rolesOk(memberRoles []string, modRoles []string) bool {
	ok := false
	for _, role := range memberRoles {
		for _, modRole := range modRoles {
			if modRole == role {
				ok = true
				break
			}
		}
	}
	return ok
}
