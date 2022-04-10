package commands

import "github.com/bwmarrin/discordgo"

func RobitLittleBro(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	bot.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{{
					Title:       "僕について",
					Description: "こんにちは！僕はロビットの弟です！マイクラ部のサーバーを管理します！",
					Color:       0xea6752,
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL:    "https://github.com/takara2314/maikurabu-robit/raw/main/secondary/robit-little-bro.png",
						Width:  128,
						Height: 128,
					},
					Author: &discordgo.MessageEmbedAuthor{
						URL:     "https://github.com/takara2314/maikurabu-robit",
						Name:    "ソースコードを見る",
						IconURL: "https://github.com/takara2314/maikurabu-robit/raw/main/secondary/robit-little-bro.png",
					},
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "/enable-server-chat",
							Value: "マイクラ鯖のチャットを確認するチャンネルにアクセスできるようになります。",
						},
						{
							Name:  "/disable-server-chat",
							Value: "マイクラ鯖のチャットを確認するチャンネルにアクセスできないようになります。",
						},
						{
							Name:  "生みの父（管理者）",
							Value: "たからーん (@takara2314)",
						},
						{
							Name:  "バージョン",
							Value: "1.4",
						},
						{
							Name:  "開発言語",
							Value: "Go",
						},
					},
				}},
			},
		},
	)
}
