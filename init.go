package main

import "time"

var (
	stopBot = make(chan bool)
)

type serverStatus struct {
	Version string
	Player  int
	Max     int
	Players []string
	Ping    time.Duration
	Icon    []byte
}
