package plugins

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cypher-org/cypher/plugins/common"
)

func Handler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	common.Handler(session, interaction)
}
