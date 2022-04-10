package commands

import "github.com/bwmarrin/discordgo"

func Start(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	bot.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "サーバーを起動しますよ！",
			},
		},
	)
}
