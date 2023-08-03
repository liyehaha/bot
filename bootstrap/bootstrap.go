package bootstrap

import (
	"log"

	"github.com/eatmoreapple/openwechat"
	"github.com/liyehaha/bot/handlers"
)

func Run() {
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop)


	log.Println("register message handler")
	bot.MessageHandler = handlers.Handler

	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl


	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")

	err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Printf("login error: %v \n", err)
			return
		}
	}
	
	bot.Block()
}
