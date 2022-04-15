package common

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type Robit struct {
	Version             string
	Primary             *RobitSession
	Secondary           *RobitSession
	Start               *StartProcess
	MaxClassmateNum     int
	ActivityInterval    time.Duration
	AutoClosingWaitTime time.Duration
	StartLocked         bool
}

type RobitSession struct {
	AppID     string
	Token     string
	Conn      *discordgo.Session
	SCommands []*discordgo.ApplicationCommand
	Stop      chan bool
}

type StartProcess struct {
	Voting              bool
	Launching           bool
	ForceRebooting      bool
	MinVoter            int
	VotePeriod          time.Duration
	ReactionCheckPeriod time.Duration
	StopVote            chan bool
}

type ServerStatus struct {
	Version string
	Player  int
	Max     int
	Players []string
	Ping    time.Duration
	Icon    []byte
}

var (
	RobitState Robit
)
