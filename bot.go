package main

import (
	"fmt"
	"log"
	"maikurabu-robit/processes"
	"os"

	"github.com/bwmarrin/discordgo"
)

func bot() {
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

	// サーバーに誰も入っていないかどうかを3分毎に確かめる
	go processes.CheckEmpty(discord)
	// ゲームステータスを1分毎に更新
	go processes.SetGameActivity(discord)

	fmt.Println("Prepare OK")
	// stopBotチャネルから何か帰ってきたら処理終了
	<-stopBot
}
