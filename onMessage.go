package main

import (
	"github.com/bwmarrin/discordgo"
)

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Deprecated

	// if strings.HasPrefix(m.Content, "/robit") {
	// 	err := commands.Robit(s, m)
	// 	if err != nil {
	// 		log.Println(err)
	// 		panic(err)
	// 	}

	// } else if strings.HasPrefix(m.Content, "/status") {
	// 	err := commands.Status(s, m)
	// 	if err != nil {
	// 		log.Println(err)
	// 		panic(err)
	// 	}

	// } else if strings.HasPrefix(m.Content, "/start") {
	// 	err := commands.Start(s, m, &isLock)
	// 	if err != nil {
	// 		log.Println(err)
	// 		panic(err)
	// 	}

	// } else if strings.HasPrefix(m.Content, "/reactions") {
	// 	splited := strings.Split(m.Content, " ")

	// 	if len(splited) == 2 {
	// 		err := commands.Reactions(s, m, splited[1])
	// 		if err != nil {
	// 			log.Println(err)
	// 			panic(err)
	// 		}
	// 	}

	// } else if strings.HasPrefix(m.Content, "/put-reaction") {
	// 	splited := strings.Split(m.Content, " ")

	// 	if len(splited) == 2 {
	// 		err := commands.PutReactions(s, m, splited[1])
	// 		if err != nil {
	// 			log.Println(err)
	// 			panic(err)
	// 		}
	// 	}

	// } else if strings.HasPrefix(m.Content, "/aed") {
	// 	err := commands.Aed(s, m, &isLock)
	// 	if err != nil {
	// 		log.Println(err)
	// 		panic(err)
	// 	}

	// } else if strings.HasPrefix(m.Content, "/stop-robit") &&
	// 	m.Author.ID == "226453185613660160" {
	// 	commands.Stop()

	// } else if strings.HasPrefix(m.Content, "/lock") &&
	// 	m.Author.ID == "226453185613660160" {
	// 	commands.Lock(s, m, m.Content, &isLock)
	// }
}
