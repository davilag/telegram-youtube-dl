package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/davilag/telego"
	"github.com/davilag/telego/api"
	"github.com/davilag/telegram-twitter-dl/mediadl"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("TELEGRAM_ACCESS_TOKEN"))
	bot := telego.Initialise(os.Getenv("TELEGRAM_ACCESS_TOKEN"))

	bot.AddCommandHandlder("download", commandHandler)
	log.Println("Listening...")
	bot.Listen()
}

func commandHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	twitterLink := strings.TrimPrefix(u.Message.Text, "/download ")
	fileName := strconv.Itoa(u.UpdateID)
	file, err := mediadl.DownloadMedia(twitterLink)
	if err != nil {
		c.SendMessage("Couldn't download the requested file")
		return nil
	}

	c.SendVideo(fileName, file)
	return nil
}
