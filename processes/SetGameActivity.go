package processes

import (
	"fmt"
	"log"
	"maikurabu-robit/types"
	"time"

	"github.com/bwmarrin/discordgo"
)

func SetGameActivity(s *discordgo.Session, isLock *bool) {
	for {
		var activityText string = "サーバーは閉まっています :("

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

		if err == nil && pcStatus != "TERMINATED" {
			if status.Player > 0 {
				activityText = fmt.Sprintf("%d人がログイン中だよ :)", status.Player)
			} else {
				activityText = "誰もログインしていないよ :|"
			}
		}

		// 操作ロックがかかっていたら
		if *isLock {
			activityText = "🔧" + activityText
		}

		// ゲームアクティビティを更新
		err = s.UpdateGameStatus(0, activityText)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		time.Sleep(30 * time.Second)
	}
}
