package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Robit(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "ロビットです :)")

	if err != nil {
		return err
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
			Name:    "ソースコードを見る",
			IconURL: "https://github.com/takara2314/maikurabu-robit/raw/main/robit.png",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "/status",
				Value: "現在のサーバーの情報を表示します。",
			},
			{
				Name:  "/start",
				Value: "サーバーを開く投票を行います。必要票数以上であれば、サーバーを開放します。",
			},
			{
				Name:  "/aed",
				Value: "サーバーの状態を分析し、必要なら強制再起動します。",
			},
			{
				Name:  "開発兼 管理者",
				Value: "たからーん (@takara2314)",
			},
			{
				Name:  "バージョン",
				Value: "1.2.1",
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
		return err
	}

	return nil
}
