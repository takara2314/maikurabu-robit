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

func checkEmpty(s *discordgo.Session) {
	ticker := time.NewTicker(common.RobitState.AutoClosingWaitTime / 5)
	defer ticker.Stop()

	const limit = 5
	var ticks int = 0
	var start int = 0
	var end int = 0
	var playerCounter []int = make([]int, limit)

	for {
		status, err := common.GetServerStatus(
			"takaran-server",
			"asia-northeast2-c",
			"minecraft-v2",
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		var mcServer *types.ServerStatus

		// Check server status if the server is not closed
		if status != "TERMINATED" {
			mcServer, err = common.GetMCServerStatus(os.Getenv("IP_ADDRESS"), 25565, time.Duration(20*time.Second))
		} else {
			ticks, start, end = 0, 0, 0
			continue
		}
		if err != nil {
			ticks, start, end = 0, 0, 0
			continue
		}

		playerCounter[end] = mcServer.Player
		ticks += 1

		end = ticks % limit
		if ticks >= limit {
			start = end + 1
			if start == limit {
				start = 0
			}

			minUnderRequired := 0

			for i := start; i < limit; i++ {
				if playerCounter[i] <= 0 {
					minUnderRequired++
				}
			}

			if start != 0 {
				for i := 0; i <= end; i++ {
					if playerCounter[i] <= 0 {
						minUnderRequired++
					}
				}
			}

			if minUnderRequired == limit && !common.RobitState.StartLocked {
				_, err = s.ChannelMessageSend(
					os.Getenv("ANNOUNCE_CHANNEL"),
					fmt.Sprintf(
						messages.AutoStopping,
						int(common.RobitState.AutoClosingWaitTime.Minutes()),
					),
				)
				if err != nil {
					log.Println(err)
					panic(err)
				}

				err = common.StopServer(
					"takaran-server",
					"asia-northeast2-c",
					"minecraft-v2",
				)
				if err != nil {
					log.Println(err)
					panic(err)
				}
			}
		}

		<-ticker.C
	}
}
