package providers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	MainMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🔑 Войти"),
		),
	)

	AuthMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("ℹ️ Обо мне"),
			tgbotapi.NewKeyboardButton("🏢 О компании"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("📂 Проекты"),
			tgbotapi.NewKeyboardButton("🚪 Выйти"),
		),
	)
)
