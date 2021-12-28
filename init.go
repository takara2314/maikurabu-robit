package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	stopBot chan bool = make(chan bool)
	discord *discordgo.Session
	isLock  bool = false
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
