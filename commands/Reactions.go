package commands

import (
	"fmt"
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
		emojiList = append(emojiList,
			fmt.Sprintf("%s %s %s", item.Emoji.ID, item.Emoji.Name, item.Emoji.User))
	}

	fmt.Println("reactions:", emojiList)

	replyMsg := strings.Join(emojiList, ", ")
	if replyMsg == "" {
		replyMsg = "None"
	}

	_, err = s.ChannelMessageSend(m.ChannelID, replyMsg)
	if err != nil {
		return err
	}

	return nil
}
