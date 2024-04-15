package plugins

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/cypher-org/cypher/plugins/common"
)

func ApplyCommands(session *discordgo.Session) {
	log.Println("Applying the common plugin commands")

	common.ApplyCommands(session, "")

	log.Println("Applied the common plugin commands")

}
