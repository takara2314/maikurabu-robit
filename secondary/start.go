package secondary

import (
	"log"
	"maikurabu-robit/common"
	"maikurabu-robit/secondary/sc"
	"os"

	"github.com/bwmarrin/discordgo"
)

func Start(ins *common.Robit) {
	var err error

	ins.Secondary.Conn, err = discordgo.New("Bot " + os.Getenv("SECONDARY_BOT_TOKEN"))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Common handler
	ins.Secondary.Conn.AddHandler(handler)

	// Connect to Discord WebSocket
	err = ins.Secondary.Conn.Open()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer ins.Secondary.Conn.Close()

	// Set activity
	go activity(ins.Secondary.Conn)

	// Register slash commands
	sc.Register(ins.Secondary.Conn, ins.Secondary.SCommands)
	defer sc.Unregister(ins.Secondary.Conn, ins.Secondary.SCommands)

	// Slash command handlers
	ins.Secondary.Conn.AddHandler(sc.Handler)

	<-ins.Secondary.Stop
}
