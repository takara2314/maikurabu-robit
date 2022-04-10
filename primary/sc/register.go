package sc

import (
	"github.com/bwmarrin/discordgo"
)

// Register registers slash commands.
func Register(bot *discordgo.Session, sCommands []*discordgo.ApplicationCommand) error {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "start",
			Description: "マイクラサーバーを起動する投票を開きます。",
		},
		{
			Name:        "status",
			Description: "マイクラサーバーの稼働状況を確認することができます。",
		},
		{
			Name:        "robit",
			Description: "ロビットのバージョンを確認できます。",
		},
	}

	sCommands = make([]*discordgo.ApplicationCommand, len(commands))

	for index, item := range commands {
		cmd, err := bot.ApplicationCommandCreate(
			bot.State.User.ID,
			"",
			item,
		)
		if err != nil {
			return err
		}

		sCommands[index] = cmd
	}

	return nil
}
