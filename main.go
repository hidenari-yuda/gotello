package main

import (
	"gotello/app/controllers"
	"gotello/config"
	"gotello/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println(controllers.StartWebServer())
}
