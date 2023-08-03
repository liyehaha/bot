package main

import (
	"gihtub.com/liyehaha/bot/config"
	"github.com/liyehaha/bot/bootstrap"
)

func main() {
	config.InitConfig("./config/config.yaml")
	bootstrap.Run()
}
