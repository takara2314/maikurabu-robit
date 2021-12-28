package commands

import (
	"maikurabu-robit/processes"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Lock(s *discordgo.Session, m *discordgo.MessageCreate, message string, isLock *bool) error {
	var args []string = strings.Split(message, " ")

	// 返信用メッセージ
	var res string
	var resLock string = "操作ロックしたよ！これで集中して作業できるね！"
	var resUnlock string = "操作ロック解除したよ！これでみんなと接しれる！"
	var resCheckLock string = "現在操作ロックされているよ！安全だね！"
	var resCheckUnlock string = "現在操作ロックが解除されているよ… 作業するならロックしておこう！"

	if len(args) >= 2 {
		switch args[1] {
		case "on":
			res = resLock
			err := processes.Lock("on", isLock)
			if err != nil {
				return err
			}

		case "off":
			res = resUnlock
			err := processes.Lock("off", isLock)
			if err != nil {
				return err
			}

		case "check":
			if *isLock {
				res = resCheckLock
			} else {
				res = resCheckUnlock
			}

		default:
			if *isLock {
				res = resUnlock
			} else {
				res = resLock
			}
			err := processes.Lock("", isLock)
			if err != nil {
				return err
			}
		}
	} else {
		if *isLock {
			res = resUnlock
		} else {
			res = resLock
		}
		err := processes.Lock("", isLock)
		if err != nil {
			return err
		}
	}

	_, err := s.ChannelMessageSend(m.ChannelID, res)
	if err != nil {
		return err
	}

	return nil
}
