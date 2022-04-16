package commands

import (
	"log"
	"maikurabu-robit/common"
	"maikurabu-robit/messages"
	"maikurabu-robit/utils"
	"os"

	"github.com/bwmarrin/discordgo"
)

func DisableServerChat(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.GuildID != os.Getenv("GUILD_ID") {
		common.ScResponseText(bot, i, messages.CannotUseOutside)
		return
	}

	member, err := bot.GuildMember(i.GuildID, i.Member.User.ID)
	if err != nil {
		log.Println(err)
		return
	}

	// Check added watchable role
	if !utils.StrSliceContains(member.Roles, os.Getenv("WATCHABLE_ROLE")) {
		common.ScResponseText(bot, i, messages.HadNotHavePermission)
		return
	}

	// Add watchable role
	err = bot.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, os.Getenv("WATCHABLE_ROLE"))
	if err != nil {
		log.Println(err)
		return
	}

	common.ScResponseText(bot, i, messages.HideWatchChannel)
}
