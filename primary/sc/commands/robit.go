package commands

import (
	"maikurabu-robit/common"

	"github.com/bwmarrin/discordgo"
)

func Robit(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	common.ScResponseEmbed(bot, i,
		[]*discordgo.MessageEmbed{{
			Title:       "僕について",
			Description: "やぁ！僕はロビットだよ！マイクラ部のサーバーを管理するよ！",
			Color:       0x05b5e6,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    "https://github.com/takara2314/maikurabu-robit/raw/main/primary/robit.png",
				Width:  128,
				Height: 128,
			},
			Author: &discordgo.MessageEmbedAuthor{
				URL:     "https://github.com/takara2314/maikurabu-robit",
				Name:    "ソースコードを見る",
				IconURL: "https://github.com/takara2314/maikurabu-robit/raw/main/primary/robit.png",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "/start",
					Value: "マイクラサーバーを起動する投票を開きます。",
				},
				{
					Name:  "/status",
					Value: "マイクラサーバーの稼働状況を確認することができます。",
				},
				{
					Name:  "生みの親（管理者）",
					Value: "たからーん (@takara2314)",
				},
				{
					Name:  "バージョン",
					Value: common.RobitState.Version,
				},
				{
					Name:  "開発言語",
					Value: "Go",
				},
			},
		}},
	)
}
