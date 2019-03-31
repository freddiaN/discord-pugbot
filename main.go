package main

import (
	"github.com/rehtlaw/discord-pugbot/do"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	do.GetServer("de", os.Getenv("DO_KEY"))
}
