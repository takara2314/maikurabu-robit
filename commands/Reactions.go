package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Reactions(s *discordgo.Session, m *discordgo.MessageCreate, messageID string) error {
	msg, err := s.ChannelMessage(m.ChannelID, messageID)
	if err != nil {
		return err
	}

	var emojiList []string
	for _, item := range msg.Reactions {
		emojiList = append(emojiList, item.Emoji.ID)
	}

	_, err = s.ChannelMessageSend(m.ChannelID, strings.Join(emojiList, ", "))
	if err != nil {
		return err
	}

	return nil
}
