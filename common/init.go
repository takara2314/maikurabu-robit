package common

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type Robit struct {
	Primary         *RobitSession
	Secondary       *RobitSession
	Start           *StartProcess
	MaxClassmateNum int
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

var (
	RobitState Robit
)
