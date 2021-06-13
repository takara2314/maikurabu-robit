package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "/robit" {
		_, err := s.ChannelMessageSend(m.ChannelID, "ロビットです :)")

		if err != nil {
			log.Println(err)
			panic(err)
		}

		embed := discordgo.MessageEmbed{
			Title:       "僕について",
			Description: "やぁ！僕はロビットだよ！マイクラ部のサーバーを管理するよ！",
			Color:       0x05b5e6,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    "https://github.com/takara2314/maikurabu-robit/raw/main/robit.png",
				Width:  128,
				Height: 128,
			},
			Author: &discordgo.MessageEmbedAuthor{
				URL:     "https://github.com/takara2314/maikurabu-robit",
				Name:    "ソースコード",
				IconURL: "https://github.com/takara2314/maikurabu-robit/raw/main/robit.png",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "/status",
					Value: "現在のサーバーの情報を表示します。",
				},
				{
					Name:  "開発者 / 管理者",
					Value: "たからーん (@takara2314)",
				},
				{
					Name:  "開発言語",
					Value: "Go",
				},
			},
		}

		_, err = s.ChannelMessageSendEmbed(
			m.ChannelID,
			&embed,
		)

		if err != nil {
			log.Println(err)
			panic(err)
		}
	}

	if m.Content == "/status" {
		_, err := s.ChannelMessageSend(m.ChannelID, "現在サーバー情報を取得しているよ！ちょっと待ってね！")

		if err != nil {
			log.Println(err)
			panic(err)
		}

		var embed discordgo.MessageEmbed
		var now time.Time = time.Now()

		status, err := getServerStatus("mc.2314.tk", 25565)
		if err != nil {
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

		} else {
			embed = discordgo.MessageEmbed{
				Title:       "サーバーの情報",
				Description: "入れます :)",
				Color:       0x34d399,
				Footer: &discordgo.MessageEmbedFooter{
					Text: fmt.Sprintf("👀 %s", time.Now().Format("2006年1月2日 15時04分05秒")),
				},
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "バージョン",
						Value: status.Version,
					},
					{
						Name: "接続プレイヤー数",
						Value: fmt.Sprintf("%d / %d",
							status.Player,
							status.Max,
						),
					},
					{
						Name:  "接続プレイヤー",
						Value: strings.Join(status.Players, ", "),
					},
					{
						Name: "遅延 (Ping)",
						Value: fmt.Sprintf("%d.%d ms",
							status.Ping.Milliseconds(),
							status.Ping.Microseconds(),
						),
					},
				},
			}
		}

		_, err = s.ChannelMessageSendEmbed(
			m.ChannelID,
			&embed,
		)

		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}
