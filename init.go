package main

import "time"

var (
	stopBot          chan bool = make(chan bool)
	isAed            bool
	isForceRebooting bool
)

type serverStatus struct {
	Version string
	Player  int
	Max     int
	Players []string
	Ping    time.Duration
	Icon    []byte
}

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	time.Local = loc

	// ボットシステムを稼働
	go bot()
}
