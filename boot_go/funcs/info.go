package funcs

import (
	"fmt"

	models "../models"
	provider "../providers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendAboutMeMenu(chatID int64, user *models.User) {
	info := fmt.Sprintf("Информация о вас:\nИмя: %s\nЛогин: %s\nКомпания: %s",
		user.Name, user.Username, user.Company)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Контактные данные", "show_contact"),
			tgbotapi.NewInlineKeyboardButtonData("Должность", "show_position"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Уровень доступа", "show_access"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, info)
	msg.ReplyMarkup = inlineKeyboard
	provider.Bot.Send(msg)
}

func HandleCallback(callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID
	data := callback.Data

	switch data {
	case "show_contact":
		SendMessage(chatID, "Контактные данные:\nEmail: user@example.com\nТелефон: +7 (123) 456-78-90")
	case "show_position":
		SendMessage(chatID, "Ваша должность: Старший разработчик")
	case "show_access":
		SendMessage(chatID, "Уровень доступа: Администратор")
	default:
		SendMessage(chatID, "Неизвестная команда")
	}

	// Ответим на callback
	provider.Bot.Send(tgbotapi.NewCallback(callback.ID, ""))
}
