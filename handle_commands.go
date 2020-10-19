package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DiscordClocks/application"
	"github.com/karashiiro/DiscordClocks/commands"
)

// CreateMessageHandler curries the message creation delegate with the provided application resources.
func CreateMessageHandler(resources *application.Resources) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(client *discordgo.Session, message *discordgo.MessageCreate) {
		messageCreateInternal(client, message, resources)
	}
}

func messageCreateInternal(client *discordgo.Session, message *discordgo.MessageCreate, resources *application.Resources) {
	if message.Author.Bot {
		return
	}

	// Command prefix
	if message.Content[0] != '^' {
		return
	}

	content := message.Content[1:]

	if strings.HasPrefix(content, "addclock") {
		commands.AddClock(resources)
	} else if strings.HasPrefix(content, "removeclock") {
		commands.RemoveClock(resources)
	}
}
