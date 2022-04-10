package sc

import (
	"maikurabu-robit/secondary/sc/commands"

	"github.com/bwmarrin/discordgo"
)

type SCHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)

func Handler(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	handlers := SCHandlers{
		"robit-little-bro":    commands.RobitLittleBro,
		"enable-server-chat":  commands.EnableServerChat,
		"disable-server-chat": commands.DisableServerChat,
	}

	if handler, ok := handlers[i.ApplicationCommandData().Name]; ok {
		handler(bot, i)
	}
}
