package commands

import (
	"maikurabu-robit/utils"
	"os"

	"github.com/bwmarrin/discordgo"
)

func EnableServerChat(bot *discordgo.Session, i *discordgo.InteractionCreate) {
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

	// Check already added watchable role
	if utils.StrSliceContains(member.Roles, os.Getenv("WATCHABLE_ROLE")) {
		bot.InteractionRespond(
			i.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "既にあなたは見る権限を持っています。",
				},
			},
		)
		return
	}

	// Add watchable role
	err = bot.GuildMemberRoleAdd(i.GuildID, i.User.ID, os.Getenv("WATCHABLE_ROLE"))
	if err != nil {
		panic(err)
	}

	bot.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "<#962246919122452531> でマイクラ鯖のチャットを見れるようになりました！チャンネルをミュートにしておくことを推奨します。",
			},
		},
	)
}
