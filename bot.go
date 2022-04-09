package main

import (
	"fmt"
	"log"
	"maikurabu-robit/processes"
	"os"

	"github.com/bwmarrin/discordgo"
)

func bot() {
	var err error

	discord, err = discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// ハンドラー
	discord.AddHandler(onMessage)

	// WebSocket開始
	err = discord.Open()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer discord.Close()

	// サーバーに誰も入っていないかどうかを3分毎に確かめる
	go processes.CheckEmpty(discord, &isLock)
	// ゲームステータスを1分毎に更新
	go processes.SetGameActivity(discord, &isLock)

	fmt.Println("Prepare OK")
	// stopBotチャネルから何か帰ってきたら処理終了
	<-stopBot
}
