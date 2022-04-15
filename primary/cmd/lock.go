package cmd

import (
	"maikurabu-robit/common"
	"maikurabu-robit/messages"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Lock(s *discordgo.Session, m *discordgo.MessageCreate, message string) error {
	var args []string = strings.Split(message, " ")
	var res string

	if len(args) >= 2 {
		switch args[1] {
		case "on":
			res = messages.StartCMDLocked
			err := common.Lock("on")
			if err != nil {
				return err
			}

		case "off":
			res = messages.StartCMDUnlocked
			err := common.Lock("off")
			if err != nil {
				return err
			}

		case "check":
			if common.RobitState.StartLocked {
				res = messages.StartCMDLocking
			} else {
				res = messages.StartCMDUnlocking
			}

		default:
			if common.RobitState.StartLocked {
				res = messages.StartCMDUnlocked
			} else {
				res = messages.StartCMDLocked
			}
			err := common.Lock("")
			if err != nil {
				return err
			}
		}

	} else {
		if common.RobitState.StartLocked {
			res = messages.StartCMDUnlocked
		} else {
			res = messages.StartCMDLocked
		}
		err := common.Lock("")
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
