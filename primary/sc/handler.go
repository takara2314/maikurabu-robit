package sc

import (
	"maikurabu-robit/primary/sc/commands"

	"github.com/bwmarrin/discordgo"
)

type SCHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)

func Handler(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	handlers := SCHandlers{
		"start":  commands.Start,
		"status": commands.Status,
		"robit":  commands.Robit,
	}

	if handler, ok := handlers[i.ApplicationCommandData().Name]; ok {
		handler(bot, i)
	}
}
