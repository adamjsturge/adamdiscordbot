package main

import "github.com/bwmarrin/discordgo"

func main() {
	_, err := discordgo.New("Bot " + "authentication token")
	if err != nil {
		panic(err)
	}

}
