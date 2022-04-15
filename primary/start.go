package primary

import (
	"log"
	"maikurabu-robit/common"
	"maikurabu-robit/primary/sc"
	"os"

	"github.com/bwmarrin/discordgo"
)

func Start() {
	var err error
	robit := common.RobitState.Primary

	robit.Conn, err = discordgo.New("Bot " + os.Getenv("PRIMARY_BOT_TOKEN"))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Common handler
	robit.Conn.AddHandler(handler)

	// Connect to Discord WebSocket
	err = robit.Conn.Open()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer robit.Conn.Close()

	// Set activity
	go activity(robit.Conn)

	// Register slash commands
	sc.Register(robit.Conn, robit.SCommands)
	defer sc.Unregister(robit.Conn, robit.SCommands)

	// Slash command handlers
	robit.Conn.AddHandler(sc.Handler)

	<-robit.Stop
}