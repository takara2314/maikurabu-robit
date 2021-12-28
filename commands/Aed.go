package commands

import (
	"maikurabu-robit/processes"

	"github.com/bwmarrin/discordgo"
)

func Aed(s *discordgo.Session, m *discordgo.MessageCreate, isLock *bool) error {
	var resMessage string

	if *isLock {
		_, err := s.ChannelMessageSend(m.ChannelID, "ごめんね。操作ロックがかけられているよ！もう少し待ったら解除されるかも…！")
		if err != nil {
			return err
		}
		return nil
	}

	if isAed && isForceRebooting {
		resMessage = "今強制再起動処理を行っているよ！しばらく待ってね！"
	} else if isAed {
		resMessage = "今解析中だよ！落ち着いて待っててね！"
	} else {
		resMessage = "強制再起動が必要かどうかを解析するよ！ちょっと待ってね！"
	}

	_, err := s.ChannelMessageSend(m.ChannelID, resMessage)
	if err != nil {
		return err
	}

	if isAed {
		return nil
	} else {
		isAed = true
	}

	pcStatus, err := processes.CheckServer()
	if err != nil {
		return err
	}

	switch pcStatus {
	case "TERMINATED":
		resMessage = "そもそもサーバー機が停止中だね。<@226453185613660160> に言って開けてもらおう！"
	default:
		resMessage = "現在たからーんの方でも対応しているみたい！ちょっとまってね！"
	}

	if pcStatus != "RUNNING" {
		_, err = s.ChannelMessageSend(m.ChannelID, resMessage)
		if err != nil {
			return err
		}

		isAed = false
		return nil
	}

	_, err = processes.GetServerStatus("mc.2314.tk", 25565)
	if err == nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "再起動は必要ないみたいだね！何か問題がある場合は、たからーんに言ってね！")
		if err != nil {
			return err
		}

		isAed = false
		return nil
	}

	_, err = s.ChannelMessageSend(m.ChannelID, "強制再起動が必要だね。5分待ってね！")
	if err != nil {
		return err
	}
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
	isAed = false

	return nil
}
