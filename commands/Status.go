package commands

import (
	"fmt"
	"maikurabu-robit/processes"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Status(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "現在サーバー情報を取得しているよ！ちょっと待ってね！")

	if err != nil {
		return err
	}

	var embed discordgo.MessageEmbed
	var now time.Time = time.Now()

	status, err := processes.GetServerStatus("mc.2314.tk", 25565)
	if err != nil {
		embed = discordgo.MessageEmbed{
			Title:       "サーバーの情報",
			Description: "閉じられています :(",
			Color:       0xdc2626,
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("👀 %s", time.Now().Format("2006年1月2日 15時04分05秒")),
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name: "検証時間",
					Value: fmt.Sprintf("%f s",
						time.Since(now).Seconds(),
					),
				},
			},
		}

	} else {
		embed = discordgo.MessageEmbed{
			Title:       "サーバーの情報",
			Description: "入れます :)",
			Color:       0x34d399,
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("👀 %s", time.Now().Format("2006年1月2日 15時04分05秒")),
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "バージョン",
					Value: status.Version,
				},
				{
					Name: "接続プレイヤー数",
					Value: fmt.Sprintf("%d / %d",
						status.Player,
						status.Max,
					),
				},
				{
					Name:  "接続プレイヤー",
					Value: strings.Join(status.Players, ", "),
				},
				{
					Name: "遅延 (Ping)",
					Value: fmt.Sprintf("%d.%d ms",
						status.Ping.Milliseconds(),
						status.Ping.Microseconds(),
					),
				},
			},
		}
	}

	_, err = s.ChannelMessageSendEmbed(
		m.ChannelID,
		&embed,
	)

	if err != nil {
		return err
	}

	return nil
}
