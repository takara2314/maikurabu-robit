package commands

import (
	"github.com/bwmarrin/discordgo"
)

func PutReactions(s *discordgo.Session, m *discordgo.MessageCreate, messageID string) error {
	err := s.MessageReactionAdd(m.ChannelID, messageID, "🤖")
	if err != nil {
		return err
	}

	return nil
}
