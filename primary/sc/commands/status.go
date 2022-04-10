package commands

import "github.com/bwmarrin/discordgo"

func Status(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	bot.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "マイクラサーバーの稼働状況を確認しますよ！",
			},
		},
	)
}
