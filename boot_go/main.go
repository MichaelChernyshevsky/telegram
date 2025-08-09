package main

import (
	"log"

	funcs "./funcs"
	provider "./providers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	var err error
	provider.Bot, err = tgbotapi.NewBotAPI("7548214792:AAGVP5ZiyUPtuuutLJEBiR-CKWZF9zjOzW8")
	if err != nil {
		log.Panic(err)
	}

	provider.Bot.Debug = true
	log.Printf("Авторизован как %s", provider.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := provider.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			if update.CallbackQuery != nil {
				funcs.HandleCallback(update.CallbackQuery)
			}
			continue
		}

		chatID := update.Message.Chat.ID

		if update.Message.IsCommand() {
			funcs.HandleCommand(update.Message)
			continue
		}

		if provider.LoginState[chatID] {
			funcs.ProcessLogin(chatID, update.Message.Text)
			provider.LoginState[chatID] = false
			continue
		}

		if _, authorized := provider.AuthorizedUsers[chatID]; !authorized {
			funcs.HandleUnauthorized(update.Message)
		} else {
			funcs.HandleAuthorized(update.Message)
		}
	}
}
