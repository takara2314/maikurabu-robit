package secondary

import (
	"fmt"
	"log"
	"maikurabu-robit/common"
	"maikurabu-robit/messages"
	"maikurabu-robit/types"
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
			mcServer, err = common.GetMCServerStatus("mc.2314.tk", 25565, time.Duration(10*time.Second))
		}

		// Show ping num info
		if err == nil && status != "TERMINATED" {
			if err == messages.ErrTimeout {
				text = messages.Timeouted
			} else if mcServer.Ping.Milliseconds() >= 100 {
				text = fmt.Sprintf(messages.HighPing, mcServer.Ping.Milliseconds())
			} else {
				text = fmt.Sprintf(messages.LowPing, mcServer.Ping.Milliseconds())
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
