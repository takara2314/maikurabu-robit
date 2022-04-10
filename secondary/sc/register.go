package sc

import (
	"github.com/bwmarrin/discordgo"
)

// Register registers slash commands.
func Register(bot *discordgo.Session, sCommands []*discordgo.ApplicationCommand) error {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "enable-server-chat",
			Description: "マイクラ鯖のチャットを確認するチャンネルにアクセスできるようになります。",
		},
		{
			Name:        "disable-server-chat",
			Description: "マイクラ鯖のチャットを確認するチャンネルにアクセスできないようになります。",
		},
		{
			Name:        "robit-little-bro",
			Description: "ロビットの弟のバージョンを確認できます。",
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
