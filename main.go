package main

import (
	"fmt"
	"os"

	"github.com/davilag/telego"
	"github.com/davilag/telego/api"
)

func main() {
	bot := telego.Initialise(os.Getenv("TELEGRAM_ACCES_TOKEN"))

	bot.AddCommandHandlder("test", commandHandler)
	fmt.Println("Listenting")
	bot.Listen()
}

func commandHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	c.SendMessage("Not sure what to do...")
	return nil
}
