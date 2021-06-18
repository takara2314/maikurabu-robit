package main

import (
	"maikurabu-robit/commands"

	"github.com/bwmarrin/discordgo"
)

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if m.Content == "/robit" {
		err := commands.Robit(s, m)
		if err != nil {
			return err
		}
		return nil

	} else if m.Content == "/status" {
		err := commands.Status(s, m)
		if err != nil {
			return err
		}
		return nil

	} else if m.Content == "/aed" {
		err := commands.Aed(s, m)
		if err != nil {
			return err
		}
		return nil

	} else if m.Content == "/stop" && m.Author.ID == "226453185613660160" {
		commands.Stop()
	}

	return nil
}
