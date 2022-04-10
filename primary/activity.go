package primary

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func activity(s *discordgo.Session) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		text := "サーバーは準備中です :("

		// Update game activity
		err := s.UpdateGameStatus(0, text)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		<-ticker.C
	}
}
