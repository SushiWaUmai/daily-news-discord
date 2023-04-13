package env

import (
	"os"
	"strconv"
)

var (
	BOT_TOKEN   string

	DB_HOST     string
	DB_NAME     string
	DB_PORT     int
	DB_USERNAME string
	DB_PASSWORD string
)

func loadEnv() {
	var err error
	BOT_TOKEN = os.Getenv("BOT_TOKEN")


	DB_HOST = os.Getenv("DB_HOST")
	if len(DB_HOST) == 0 {
		DB_HOST = "localhost"
	}

	DB_NAME = os.Getenv("DB_NAME")
	if len(DB_NAME) == 0 {
		DB_NAME = "DailyNewsDiscord"
	}

	DB_PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		DB_PORT = 5432
	}

	DB_USERNAME = os.Getenv("DB_USERNAME")
	if len(DB_USERNAME) == 0 {
		DB_USERNAME = "postgres"
	}

	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	if len(DB_PASSWORD) == 0 {
		DB_PASSWORD = "postgres"
	}
}
