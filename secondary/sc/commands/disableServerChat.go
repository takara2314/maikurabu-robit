package commands

import (
	"maikurabu-robit/utils"
	"os"

	"github.com/bwmarrin/discordgo"
)

func DisableServerChat(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.GuildID != os.Getenv("GUILD_ID") {
		bot.InteractionRespond(
			i.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "このコマンドは、マイクラ部以外のサーバーでは使用できません。",
				},
			},
		)
		return
	}

	member, err := bot.GuildMember(i.GuildID, i.User.ID)
	if err != nil {
		panic(err)
	}

	// Check added watchable role
	if !utils.StrSliceContains(member.Roles, os.Getenv("WATCHABLE_ROLE")) {
		bot.InteractionRespond(
			i.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "初めから見る権限を持っていません。",
				},
			},
		)
		return
	}

	// Add watchable role
	err = bot.GuildMemberRoleRemove(i.GuildID, i.User.ID, os.Getenv("WATCHABLE_ROLE"))
	if err != nil {
		panic(err)
	}

	bot.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "マイクラ鯖のチャットを見れないようにしました。",
			},
		},
	)
}
