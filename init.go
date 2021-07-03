package main

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

var (
	stopBot chan bool = make(chan bool)
	discord *discordgo.Session
)

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	time.Local = loc

	// ボットシステムを稼働
	go bot()
}
