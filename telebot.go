package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {

	botToken := "6961681806:AAHajJKNBG_s9Muzt6EqgRV2pXAfWSR9k1E"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID := tu.ID(update.Message.Chat.ID)

			_, _ = bot.CopyMessage(
				tu.CopyMessage(
					chatID,
					chatID,
					update.Message.MessageID,
				),
			)
		}
	}
}
