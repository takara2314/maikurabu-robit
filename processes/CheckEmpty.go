package processes

import (
	"log"
	"maikurabu-robit/types"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func CheckEmpty(s *discordgo.Session) {
	const limit = 5
	var ticks int = 0
	var start int = 0
	var end int = 0
	var playerCounter []int = make([]int, limit)

	for {
		pcStatus, err := CheckServer()
		if err != nil {
			log.Println(err)
			panic(err)
		}

		var status *types.ServerStatus

		// サーバー機が閉じられていなかったら確認
		if pcStatus != "TERMINATED" {
			status, err = GetServerStatus("mc.2314.tk", 25565)
		}
		if err != nil {
			continue
		}

		playerCounter[end] = status.Player
		ticks += 1

		end = ticks % limit
		if ticks >= limit {
			start = end + 1
			if start == limit {
				start = 0
			}

			minUnderRequired := 0

			for i := start; i < limit; i++ {
				if playerCounter[i] < 0 {
					minUnderRequired++
				}
			}

			if start != 0 {
				for i := 0; i <= end; i++ {
					if playerCounter[i] < 0 {
						minUnderRequired++
					}
				}
			}

			if minUnderRequired == limit {
				_, err = s.ChannelMessageSend(
					os.Getenv("ANNOUNCE_CHANNEL_ID"),
					"15分間サーバーに誰もいないので、サーバーを終了するよ。",
				)
				if err != nil {
					log.Println(err)
					panic(err)
				}

				err = StopServer()
				if err != nil {
					log.Println(err)
					panic(err)
				}
			}
		}

		time.Sleep(3 * time.Minute)
	}
}
