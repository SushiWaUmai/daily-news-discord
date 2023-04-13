package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SushiWaUmai/daily-news-discord/bot"
	"github.com/SushiWaUmai/daily-news-discord/env"
)

func main() {
	news, err := bot.CreateBot(env.BOT_TOKEN)

	if err != nil {
		log.Panicln("Failed to create discord bot", err)
	}

	err = news.Start()
	defer news.Close()

	if err != nil {
		log.Panicln("Error while running discord bot", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-c
}
