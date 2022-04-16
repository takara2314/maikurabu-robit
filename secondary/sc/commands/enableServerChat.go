package commands

import (
	"fmt"
	"log"
	"maikurabu-robit/common"
	"maikurabu-robit/messages"
	"maikurabu-robit/utils"
	"os"

	"github.com/bwmarrin/discordgo"
)

func EnableServerChat(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.GuildID != os.Getenv("GUILD_ID") {
		common.ScResponseText(bot, i, messages.CannotUseOutside)
		return
	}

	member, err := bot.GuildMember(i.GuildID, i.Member.User.ID)
	if err != nil {
		log.Println(err)
		return
	}

	// Check already added watchable role
	if utils.StrSliceContains(member.Roles, os.Getenv("WATCHABLE_ROLE")) {
		common.ScResponseText(bot, i, messages.AlreadyHavePermission)
		return
	}

	// Add watchable role
	err = bot.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, os.Getenv("WATCHABLE_ROLE"))
	if err != nil {
		log.Println(err)
		return
	}

	common.ScResponseText(bot, i,
		fmt.Sprintf(messages.ShowWatchChannel, os.Getenv("WATCH_CHANNEL")),
	)
}
