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
)

func main() {
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("TELEGRAM_ACCESS_TOKEN"))
	bot := telego.Initialise(os.Getenv("TELEGRAM_ACCESS_TOKEN"))

	bot.AddCommandHandler("download", commandHandler)
	bot.AddCommandHandler("help", helpHandler)
	bot.SetDefaultMessageHandler(helpHandler)
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

func helpHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	c.SendMessage(`This bot downloads videos from different platforms using youtube-dl. 
	Commands:
	  - /download [url] - Downloads the video given the specific URL. The video has to be smaller than 50Mb
	`)
	return nil
}
