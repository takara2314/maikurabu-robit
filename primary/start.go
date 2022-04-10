package primary

import (
	"log"
	"maikurabu-robit/common"
	"maikurabu-robit/primary/sc"
	"os"

	"github.com/bwmarrin/discordgo"
)

func Start(ins *common.Robit) {
	var err error

	ins.Primary.Conn, err = discordgo.New("Bot " + os.Getenv("PRIMARY_BOT_TOKEN"))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Common handler
	ins.Primary.Conn.AddHandler(handler)

	// Connect to Discord WebSocket
	err = ins.Primary.Conn.Open()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer ins.Primary.Conn.Close()

	// Set activity
	go activity(ins.Primary.Conn)

	// Register slash commands
	sc.Register(ins.Primary.Conn, ins.Primary.SCommands)
	defer sc.Unregister(ins.Primary.Conn, ins.Primary.SCommands)

	// Slash command handlers
	ins.Primary.Conn.AddHandler(sc.Handler)

	<-ins.Primary.Stop
}
