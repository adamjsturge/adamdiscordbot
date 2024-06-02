package main

import (
	"flag"
	"fmt"

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
	// discord.AddHandler(discordPrefixedCommands)

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"helloworld": basicCommand,
	}

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func basicCommand(discord *discordgo.Session, i *discordgo.InteractionCreate) {
	fmt.Println("Command executed")
	discord.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Congratulations, you just executed your first slash command",
		},
	})
	// discord.ChannelMessageSend(message.ChannelID, "Hello, world!")
}

// func discordPrefixedCommands(discord *discordgo.Session, message *discordgo.MessageCreate) {
// 	if message.Content[:1] != commandPrefix || message.Content == "" {
// 		return
// 	}

// 	switch message.Content[1:] {
// 	case "helloworld":
// 		discord.ChannelMessageSend(message.ChannelID, "Hello, world!")
// 	default:
// 		discord.ChannelMessageSend(message.ChannelID, "Unknown command")
// 	}
// }
