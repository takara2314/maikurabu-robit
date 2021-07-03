package processes

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

func BoardCast(s *discordgo.Session, text string) error {
	_, err := s.ChannelMessageSend(
		os.Getenv("ANNOUNCE_CHANNEL_ID"),
		text,
	)
	if err != nil {
		return err
	}

	return nil
}
