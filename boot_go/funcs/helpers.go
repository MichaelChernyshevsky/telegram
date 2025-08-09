package funcs

import (
	provider "../providers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendMainMenu(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = provider.MainMenuKeyboard
	provider.Bot.Send(msg)
}

func SendAuthMenu(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = provider.AuthMenuKeyboard
	provider.Bot.Send(msg)
}

func SendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	provider.Bot.Send(msg)
}
