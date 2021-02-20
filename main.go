package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/davilag/telego"
	"github.com/davilag/telego/api"
	"github.com/davilag/telegram-twitter-dl/twitterdl"
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

func sendErrorMessage(c telego.Conversation) {
	c.SendMessage("Couldn't download the requested file")
}

func commandHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	twitterLink := strings.TrimPrefix(u.Message.Text, "/download ")
	fileName := strconv.Itoa(u.UpdateID)

	err := twitterdl.DownloadTwitterMedia(twitterLink, fileName)
	if err != nil {
		sendErrorMessage(c)
		return nil
	}

	defer os.Remove(fileName)
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		sendErrorMessage(c)
		return nil
	}

	c.SendVideo(content)
	return nil
}
