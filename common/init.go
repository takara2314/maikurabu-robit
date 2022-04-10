package common

import "github.com/bwmarrin/discordgo"

type Robit struct {
	Primary   *RobitSession
	Secondary *RobitSession
}

type RobitSession struct {
	Conn      *discordgo.Session
	SCommands []*discordgo.ApplicationCommand
	Stop      chan bool
}
