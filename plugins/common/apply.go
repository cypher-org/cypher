package common

import (
	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "help",
		Description: "Displays a simple help menu",
	},
	{
		Name:        "ping",
		Description: "Displays the bot's heartbeat latency",
	},
}

func ApplyCommands(session *discordgo.Session, guildID string) {
	for _, command := range Commands {
		session.ApplicationCommandCreate(session.State.User.ID, guildID, command)
	}
}
