package providers

import (
	"sync"

	"../models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	Bot             *tgbotapi.BotAPI
	AuthorizedUsers = make(map[int64]*models.User)
	UsersMutex      sync.Mutex
	LoginState      = make(map[int64]bool)
	PredefinedUsers = []models.User{
		{
			Username: "admin",
			Password: "admin123",
			Name:     "Администратор",
			Company:  "ТехноКомпания",
			Projects: []string{"Проект А", "Проект Б"},
		},
	}
)
