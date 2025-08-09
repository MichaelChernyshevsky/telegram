package funcs

import (
	"strings"

	models "../models"
	provider "../providers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleUnauthorized(msg *tgbotapi.Message) {
	switch msg.Text {
	case "🔑 Войти":
		RequestLogin(msg.Chat.ID)
	default:
		SendMainMenu(msg.Chat.ID, "Пожалуйста, сначала войдите в систему.")
	}
}

func HandleAuthorized(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	user := provider.AuthorizedUsers[chatID]

	switch msg.Text {
	case "ℹ️ Обо мне":
		SendAboutMeMenu(chatID, user)
	case "🏢 О компании":
		SendMessage(chatID, "Компания: "+user.Company+"\nКонтактные данные: ...")
	case "📂 Проекты":
		projects := "Ваши проекты:\n" + strings.Join(user.Projects, "\n")
		SendMessage(chatID, projects)
	case "🚪 Выйти":
		provider.UsersMutex.Lock()
		delete(provider.AuthorizedUsers, chatID)
		provider.UsersMutex.Unlock()
		SendMainMenu(chatID, "Вы вышли из системы.")
	default:
		SendAuthMenu(chatID, "Выберите действие:")
	}
}

func RequestLogin(chatID int64) {
	provider.LoginState[chatID] = true
	msg := tgbotapi.NewMessage(chatID, "Введите логин и пароль в формате: логин:пароль\nПример: admin:admin123")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	provider.Bot.Send(msg)
}

func ProcessLogin(chatID int64, text string) {
	credentials := strings.SplitN(text, ":", 2)
	if len(credentials) != 2 {
		SendMessage(chatID, "Неверный формат. Используйте: логин:пароль")
		return
	}

	username := strings.TrimSpace(credentials[0])
	password := strings.TrimSpace(credentials[1])

	for _, user := range provider.PredefinedUsers {
		if user.Username == username && user.Password == password {
			provider.UsersMutex.Lock()
			provider.AuthorizedUsers[chatID] = &models.User{
				Username: user.Username,
				Name:     user.Name,
				Company:  user.Company,
				Projects: user.Projects,
			}
			provider.UsersMutex.Unlock()
			SendAuthMenu(chatID, "Авторизация успешна! Добро пожаловать, "+user.Name+"!")
			return
		}
	}

	SendMessage(chatID, "Неверный логин или пароль")
	SendMainMenu(chatID, "Попробуйте еще раз или обратитесь к администратору.")
}

func HandleCommand(msg *tgbotapi.Message) {
	switch msg.Command() {
	case "start":
		SendMainMenu(msg.Chat.ID, "Добро пожаловать! Для доступа к функциям войдите в систему.")
	case "help":
		SendMessage(msg.Chat.ID, "Доступные команды:\n/start - начать работу\n/help - помощь")
	}
}
