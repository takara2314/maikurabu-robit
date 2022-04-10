package sc

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Register deletes slash commands.
func Unregister(bot *discordgo.Session, sCommands []*discordgo.ApplicationCommand) {
	for _, cmd := range sCommands {
		err := bot.ApplicationCommandDelete(
			bot.State.User.ID,
			"",
			cmd.ID,
		)
		if err != nil {
			log.Printf("Cannot delete '%v' command: %v", cmd.Name, err)
		}
	}
}
