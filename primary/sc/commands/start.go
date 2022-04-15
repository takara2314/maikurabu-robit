package commands

import (
	"fmt"
	"log"
	"maikurabu-robit/common"
	"maikurabu-robit/messages"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Start(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	startInfo := common.RobitState.Start

	// Server computer status
	status, err := common.GetServerStatus(
		"takaran-server",
		"asia-northeast2-c",
		"minecraft-v2",
	)
	if err != nil {
		log.Println(err)
		return
	}

	booted := false
	connectable := false

	if status != "RUNNING" {
		booted = true
	}

	if booted {
		_, err := common.GetMCServerStatus(os.Getenv("IP_ADDRESS"), 25565)
		if err == nil {
			connectable = true
		}
	}

	resMessage := ""
	if booted {
		if connectable {
			resMessage = messages.Connectable
		} else {
			resMessage = messages.ForceRebootByProblem
		}
	} else {
		if startInfo.Voting {
			resMessage = messages.VotingNow
		} else {
			resMessage = messages.StartVote
		}
	}

	if startInfo.Launching {
		resMessage = messages.LaunchingNow
	}

	common.ScResponseText(bot, i, resMessage)

	// If already the voting has been started or connectable or server has been launching, end to start process.
	if startInfo.Voting || connectable || startInfo.Launching {
		return
	}

	///////////////////////////////////
	//  Start voting
	///////////////////////////////////

	startInfo.Voting = true
	defer func() {
		startInfo.Voting = false
	}()

	if !booted {
		now := time.Now()

		msg, err := bot.ChannelMessageSendEmbed(
			i.ChannelID,
			&discordgo.MessageEmbed{
				Title:       "サーバーを開くかの投票",
				Description: "この投票次第でサーバーを開くかどうかが決まるよ！",
				Color:       0x4338ca,
				Footer: &discordgo.MessageEmbedFooter{
					Text: fmt.Sprintf("✋ %s", now.Format("2006年1月2日 15時04分05秒")),
				},
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "必要投票人数",
						Value: fmt.Sprintf("%d人", startInfo.MinVoter),
					},
					{
						Name: "投票期限",
						Value: fmt.Sprintf("%d分後 (%d時%d分)",
							int(startInfo.VotePeriod.Minutes()),
							now.Add(startInfo.VotePeriod).Hour(),
							now.Add(startInfo.VotePeriod).Minute(),
						),
					},
				},
			},
		)
		if err != nil {
			log.Println(err)
			return
		}

		err = bot.MessageReactionAdd(i.ChannelID, msg.ID, "👍")
		if err != nil {
			log.Println(err)
			return
		}

		///////////////////////////////////
		//  Wait finish of voting
		///////////////////////////////////

		expiredStop := make(chan bool)

		go waitForVoting(
			&expiredStop,
			&startInfo.StopVote,
			startInfo.VotePeriod,
		)

		go checkReactionForVoting(
			&expiredStop,
			&startInfo.StopVote,
			startInfo.ReactionCheckPeriod,
			bot,
			i.ChannelID,
			msg.ID,
			startInfo.MinVoter,
			common.RobitState.MaxClassmateNum,
		)

		<-startInfo.StopVote

		///////////////////////////////////
		//  Finish voting
		///////////////////////////////////

		met, err := isMeetMinVoter(
			bot,
			i.ChannelID,
			msg.ID,
			startInfo.MinVoter,
			common.RobitState.MaxClassmateNum,
		)
		if err != nil {
			log.Print(err)
			return
		}

		if !met {
			common.ScResponseText(bot, i, messages.NotMeetMinVoter)
			return
		}

		common.ScResponseText(bot, i,
			fmt.Sprintf(messages.MeetMinVoter, startInfo.MinVoter),
		)

		///////////////////////////////////
		//  Start server
		///////////////////////////////////

		startInfo.Launching = true
		defer func() {
			startInfo.Launching = false
		}()

		err = common.StartServer(
			"takaran-server",
			"asia-northeast2-c",
			"minecraft-v2",
		)
		if err != nil {
			common.ScResponseText(bot, i,
				fmt.Sprintf(messages.FailedLaunchServer, os.Getenv("ADMIN_DISCORD_ID")),
			)
		} else {
			common.ScResponseText(bot, i, messages.LaunchServer)
		}

	} else {
		startInfo.ForceRebooting = true
		defer func() {
			startInfo.ForceRebooting = false
		}()

		err = common.RebootServer(
			"takaran-server",
			"asia-northeast2-c",
			"minecraft-v2",
		)
		if err != nil {
			common.ScResponseText(bot, i,
				fmt.Sprintf(messages.FailedForceReboot, os.Getenv("ADMIN_DISCORD_ID")),
			)
		} else {
			common.ScResponseText(bot, i, messages.ForceRebooted)
		}
	}
}

func waitForVoting(expiredStop *chan bool, stopVote *chan bool, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	<-ticker.C
	*expiredStop <- true
	*stopVote <- true
}

func checkReactionForVoting(expiredStop *chan bool, stopVote *chan bool, interval time.Duration, s *discordgo.Session, chanId string, msgId string, minNum int, maxNum int) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-*expiredStop:
			return

		default:
			met, err := isMeetMinVoter(
				s,
				chanId,
				msgId,
				minNum,
				maxNum,
			)
			if err != nil {
				log.Print(err)
				return
			}

			if met {
				*stopVote <- true
			}

			<-ticker.C
		}
	}
}

func isMeetMinVoter(s *discordgo.Session, chanId string, msgId string, minNum int, maxNum int) (bool, error) {
	users, err := s.MessageReactions(
		chanId,
		msgId,
		"👍",
		maxNum,
		"",
		"",
	)
	if err != nil {
		return false, nil
	}

	if len(users) < minNum+1 {
		return false, nil
	}
	return true, nil
}
