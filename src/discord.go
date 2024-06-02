package main

import (
	"flag"

	"github.com/bwmarrin/discordgo"
)

var (
	commandPrefix = "!"
	GuildID       = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "helloworld",
			Description: "Sends a hello world message",
		},
	}
)

func discordAddHandlers(discord *discordgo.Session) {
	discord.AddHandler(discordPrefixedCommands)
}

func discordPrefixedCommands(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content[:1] != commandPrefix || message.Content == "" {
		return
	}

	switch message.Content[1:] {
	case "helloworld":
		discord.ChannelMessageSend(message.ChannelID, "Hello, world!")
	default:
		discord.ChannelMessageSend(message.ChannelID, "Unknown command")
	}
}
