package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "たからーん" || m.Content == "たからん" {
		_, err := s.ChannelMessageSend(m.ID, "なーに！")

		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}
