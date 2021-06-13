package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	stopBot = make(chan bool)
)

func main() {
	discord, err := discordgo.New()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// トークン
	discord.Token = "Bot " + os.Getenv("BOT_TOKEN")

	// ハンドラー
	discord.AddHandler(onMessage)

	// WebSocket開始
	err = discord.Open()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer discord.Close()

	fmt.Println("Prepare OK")
	// stopBotチャネルから何か帰ってきたら処理終了
	<-stopBot
}
