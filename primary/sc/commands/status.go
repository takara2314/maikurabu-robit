package commands

import (
	"fmt"
	"log"
	"maikurabu-robit/common"
	"maikurabu-robit/messages"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Status(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	var embed discordgo.MessageEmbed
	var mcServer *common.ServerStatus
	var now time.Time = time.Now()

	common.ScResponseText(bot, i, messages.CheckStatusWait)

	// Server computer status
	status, err := common.GetServerStatus(
		"takaran-server",
		"asia-northeast2-c",
		"minecraft-v2",
	)
	if err != nil {
		log.Println(err)
		return
	}

	// If server is not shutdowned, check server info
	if status != "TERMINATED" {
		mcServer, err = common.GetMCServerStatus(os.Getenv("IP_ADDRESS"), 25565, time.Duration(10*time.Second))
	}

	if err != nil || status == "TERMINATED" {
		if err == messages.ErrTimeout {
			embed = discordgo.MessageEmbed{
				Title:       "サーバーの情報",
				Description: "タイムアウトしました :(",
				Color:       0xdc7526,
				Footer: &discordgo.MessageEmbedFooter{
					Text: fmt.Sprintf("👀 %s", time.Now().Format("2006年1月2日 15時04分05秒")),
				},
				Fields: []*discordgo.MessageEmbedField{
					{
						Name: "検証時間",
						Value: fmt.Sprintf("%f s",
							time.Since(now).Seconds(),
						),
					},
				},
			}

		} else {
			embed = discordgo.MessageEmbed{
				Title:       "サーバーの情報",
				Description: "閉じられています :(",
				Color:       0xdc2626,
				Footer: &discordgo.MessageEmbedFooter{
					Text: fmt.Sprintf("👀 %s", time.Now().Format("2006年1月2日 15時04分05秒")),
				},
				Fields: []*discordgo.MessageEmbedField{
					{
						Name: "検証時間",
						Value: fmt.Sprintf("%f s",
							time.Since(now).Seconds(),
						),
					},
				},
			}
		}

	} else {
		var description string = "入れます :)"
		var color int = 0x34d399

		if mcServer.Ping.Milliseconds() >= 100 {
			description = "入れますが少しラグいです :|"
			color = 0xfbbf24
		}

		var players string = strings.Join(mcServer.Players, ", ")
		if len(mcServer.Players) == 0 {
			players = "誰もいません"
		}

		embed = discordgo.MessageEmbed{
			Title:       "サーバーの情報",
			Description: description,
			Color:       color,
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("👀 %s", time.Now().Format("2006年1月2日 15時04分05秒")),
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "バージョン",
					Value: mcServer.Version,
				},
				{
					Name: "接続プレイヤー数",
					Value: fmt.Sprintf("%d / %d",
						mcServer.Player,
						mcServer.Max,
					),
				},
				{
					Name:  "接続プレイヤー",
					Value: players,
				},
				{
					Name: "遅延 (Ping)",
					Value: fmt.Sprintf("%d.%d ms",
						mcServer.Ping.Milliseconds(),
						mcServer.Ping.Microseconds(),
					),
				},
			},
		}
	}

	_, err = bot.ChannelMessageSendEmbed(
		i.ChannelID,
		&embed,
	)

	if err != nil {
		log.Println(err)
		return
	}
}
