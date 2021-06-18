package commands

import (
	"log"
)

func Stop() {
	log.Println("コマンドによる強制終了")
	panic("force stop by user command")
}
