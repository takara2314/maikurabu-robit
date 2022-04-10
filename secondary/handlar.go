package secondary

import (
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "ロビ弟") {
		_, err := s.ChannelMessageSend(
			m.ChannelID,
			"呼んだー？",
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(m.Content, "/stop-robit") &&
		m.Author.ID == os.Getenv("ADMIN_DISCORD_ID") {

		s.ChannelMessageSend(
			m.ChannelID,
			"force stop",
		)

		panic("force stop")
	}
}
