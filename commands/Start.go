package commands

import (
	"fmt"
	"maikurabu-robit/processes"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Start(s *discordgo.Session, m *discordgo.MessageCreate) error {
	var resMessage string
	var isOpeningPC bool
	var isOpeningServer bool

	// サーバー機が開いているかをチェック
	pcStatus, err := processes.CheckServer()
	if err != nil {
		return err
	}

	switch pcStatus {
	case "RUNNING":
		isOpeningPC = true
	default:
		isOpeningPC = false
	}

	// サーバー機が開いているなら、サーバーに入れるかをチェック
	if isOpeningPC {
		_, err := processes.GetServerStatus("mc.2314.tk", 25565)
		if err == nil {
			isOpeningServer = true
		} else {
			isOpeningServer = false
		}
	}

	// もしサーバー機が開いていないなら
	if !isOpeningPC {
		// サーバーを開くかの投票が開始していないなら
		if isStartVoting {
			resMessage = "今サーバーを起動するかの投票をとっているよ！上の投稿を見てね！"
		} else {
			resMessage = "サーバーを起動してほしい人は、次の投稿に「いいね」をつけよう！"
		}

		// サーバー機は開いているが、サーバーに入れないなら
	} else {
		if isOpeningServer {
			resMessage = "今はサーバーが開いているみたいだね。参加できるよ！"
		} else {
			resMessage = "サーバーは今不調みたいだね。今から強制再起動するよ！"
		}
	}

	_, err = s.ChannelMessageSend(m.ChannelID, resMessage)
	if err != nil {
		return err
	}

	// 投票が始まっているか、サーバーに入れるなら
	if isStartVoting || isOpeningServer {
		return nil
	}

	// 投票開始
	isStartVoting = true

	// サーバー機が開いていないなら
	if !isOpeningPC {
		now := time.Now()

		embed := discordgo.MessageEmbed{
			Title:       "サーバーを開くかの投票",
			Description: "この投票次第でサーバーを開くかどうかが決まるよ！",
			Color:       0x4338ca,
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("✋ %s", now.Format("2006年1月2日 15時04分05秒")),
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "必要投票人数",
					Value: "3人",
				},
				{
					Name: "投票期限",
					Value: fmt.Sprintf("5分後 (%d時%d分)",
						now.Add(5*time.Minute).Hour(),
						now.Add(5*time.Minute).Minute(),
					),
				},
			},
		}

		_, err = s.ChannelMessageSendEmbed(
			m.ChannelID,
			&embed,
		)
		if err != nil {
			return err
		}

		// time.Sleep(5 * time.Minute)

	} else {
		isForceRebooting = true

		err = processes.RebootServer()
		if err != nil {
			resMessage = "強制再起動が失敗しちゃった… <@226453185613660160> に言ってね！"
		} else {
			resMessage = "強制再起動したよ！待たせてごめんね！"
		}

		_, err = s.ChannelMessageSend(m.ChannelID, resMessage)
		if err != nil {
			return err
		}

		isForceRebooting = false
	}

	isStartVoting = false
	return nil
}
