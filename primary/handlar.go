package primary

import (
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "ロビット") {
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

	// if strings.HasPrefix(m.Content, "/robit") {
	// 	err := commands.Robit(s, m)
	// 	if err != nil {
	// 		log.Println(err)
	// 		panic(err)
	// 	}

	// } else if strings.HasPrefix(m.Content, "/status") {
	// 	err := commands.Status(s, m)
	// 	if err != nil {
	// 		log.Println(err)
	// 		panic(err)
	// 	}

	// } else if strings.HasPrefix(m.Content, "/start") {
	// 	err := commands.Start(s, m, &isLock)
	// 	if err != nil {
	// 		log.Println(err)
	// 		panic(err)
	// 	}

	// } else if strings.HasPrefix(m.Content, "/stop-robit") &&
	// 	m.Author.ID == "226453185613660160" {
	// 	commands.Stop()

	// } else if strings.HasPrefix(m.Content, "/lock") &&
	// 	m.Author.ID == "226453185613660160" {
	// 	commands.Lock(s, m, m.Content, &isLock)
	// }
}
