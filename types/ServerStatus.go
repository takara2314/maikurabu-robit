package types

import "time"

type ServerStatus struct {
	Version string
	Player  int
	Max     int
	Players []string
	Ping    time.Duration
	Icon    []byte
}
