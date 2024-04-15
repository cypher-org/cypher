package common

import "github.com/bwmarrin/discordgo"

var CommandHandlersMap = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"help": handleHelpCommand,
}

func handleHelpCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Unimplemented yet!",
		},
	})
}

func Handler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type == discordgo.InteractionApplicationCommand {
		if handler, ok := CommandHandlersMap[interaction.ApplicationCommandData().Name]; ok {
			handler(session, interaction)
		}
	}
}
