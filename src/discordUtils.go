package main

import "github.com/bwmarrin/discordgo"

func createDiscordTextChannel(channel_name string, discord *discordgo.Session, guildID string) (string, error) {
	channel, err := discord.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
		Name: channel_name,
		Type: discordgo.ChannelTypeGuildText,
		// ParentID: "123456789012345678",
	})
	if err != nil {
		return "", err
	}
	return channel.ID, nil
}
