package common

import "github.com/bwmarrin/discordgo"

func ScResponseText(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msg,
			},
		},
	)
}

func ScFollowupText(s *discordgo.Session, appID string, i *discordgo.InteractionCreate, msg string) {
	s.FollowupMessageCreate(appID, i.Interaction, true, &discordgo.WebhookParams{
		Content: msg,
	})
}

func ScResponseEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, embeds []*discordgo.MessageEmbed) {
	s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: embeds,
			},
		},
	)
}
