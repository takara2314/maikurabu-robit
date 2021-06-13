package main

import (
	"log"

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
				Name:    "ロビット",
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

	// if strings.HasPrefix(m.Content, "/status") {
	// 	embed := discordgo.MessageEmbed{
	// 		URL:   "test",
	// 		Title: "test",
	// 	}

	// 	_, err := s.ChannelMessageSendEmbed(
	// 		m.ChannelID,
	// 		&embed,
	// 	)

	// 	if err != nil {
	// 		log.Println(err)
	// 		panic(err)
	// 	}
	// }
}
