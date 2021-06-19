package main

import (
	"log"
	"maikurabu-robit/commands"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "/robit") {
		err := commands.Robit(s, m)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(m.Content, "/status") {
		err := commands.Status(s, m)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(m.Content, "/start") {
		err := commands.Start(s, m)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(m.Content, "/reactions") {
		splited := strings.Split(m.Content, " ")

		if len(splited) == 2 {
			err := commands.Reactions(s, m, splited[1])
			if err != nil {
				log.Println(err)
				panic(err)
			}
		}

	} else if strings.HasPrefix(m.Content, "/put-reaction") {
		splited := strings.Split(m.Content, " ")

		if len(splited) == 2 {
			err := commands.PutReactions(s, m, splited[1])
			if err != nil {
				log.Println(err)
				panic(err)
			}
		}

	} else if strings.HasPrefix(m.Content, "/aed") {
		err := commands.Aed(s, m)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(m.Content, "/stop") &&
		m.Author.ID == "226453185613660160" {
		commands.Stop()
	}
}
