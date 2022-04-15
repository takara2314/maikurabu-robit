package secondary

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "/stop-robit") &&
		m.Author.ID == os.Getenv("ADMIN_DISCORD_ID") {

		s.ChannelMessageSend(
			m.ChannelID,
			"force stop",
		)

		panic("force stop")
	}
}
