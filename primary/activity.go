package primary

import (
	"fmt"
	"log"
	"maikurabu-robit/common"
	"maikurabu-robit/messages"
	"maikurabu-robit/types"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func activity(s *discordgo.Session) {
	ticker := time.NewTicker(common.RobitState.ActivityInterval)
	defer ticker.Stop()

	for {
		text := messages.Closing

		status, err := common.GetServerStatus(
			"takaran-server",
			"asia-northeast2-c",
			"minecraft-v2",
		)
		if err != nil {
			log.Println(err)
			return
		}

		var mcServer *types.ServerStatus

		// Check status if server is not closed
		if status != "TERMINATED" {
			mcServer, err = common.GetMCServerStatus(os.Getenv("IP_ADDRESS"), 25565, time.Duration(10*time.Second))
		}

		// Show player num info
		if err == nil && status != "TERMINATED" {
			if err == messages.ErrTimeout {
				text = messages.Timeouted
			} else if mcServer.Player > 0 {
				text = fmt.Sprintf(messages.HowManyPlayers, mcServer.Player)
			} else {
				text = messages.NoPlayers
			}
		}

		// If start command is locked
		if common.RobitState.StartLocked {
			text = "🔒 " + text
		}

		// Update game activity
		err = s.UpdateGameStatus(0, text)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		<-ticker.C
	}
}
