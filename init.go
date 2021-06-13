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
