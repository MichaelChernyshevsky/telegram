package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type User struct {
	Username string
	Password string
	Name     string
	Company  string
	Projects []string
}

var (
	bot             *tgbotapi.BotAPI
	authorizedUsers = make(map[int64]*User)
	usersMutex      sync.Mutex
	loginState      = make(map[int64]bool)
	predefinedUsers = []User{
		{
			Username: "admin",
			Password: "admin123",
			Name:     "Администратор",
			Company:  "ТехноКомпания",
			Projects: []string{"Проект А", "Проект Б"},
		},
	}
)

var (
	mainMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🔑 Войти"),
		),
	)
	authMenuKeyboard = tgbotapi.NewReplyKeyboard(
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

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI("7548214792:AAGVP5ZiyUPtuuutLJEBiR-CKWZF9zjOzW8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Авторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			// Обработка callback-ов от inline-кнопок
			if update.CallbackQuery != nil {
				handleCallback(update.CallbackQuery)
			}
			continue
		}

		chatID := update.Message.Chat.ID

		if update.Message.IsCommand() {
			handleCommand(update.Message)
			continue
		}

		if loginState[chatID] {
			processLogin(chatID, update.Message.Text)
			loginState[chatID] = false
			continue
		}

		if _, authorized := authorizedUsers[chatID]; !authorized {
			handleUnauthorized(update.Message)
		} else {
			handleAuthorized(update.Message)
		}
	}
}

func handleCallback(callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID
	data := callback.Data

	switch data {
	case "show_contact":
		sendMessage(chatID, "Контактные данные:\nEmail: user@example.com\nТелефон: +7 (123) 456-78-90")
	case "show_position":
		sendMessage(chatID, "Ваша должность: Старший разработчик")
	case "show_access":
		sendMessage(chatID, "Уровень доступа: Администратор")
	default:
		sendMessage(chatID, "Неизвестная команда")
	}

	// Ответим на callback, чтобы убрать "часики" на кнопке
	bot.Send(tgbotapi.NewCallback(callback.ID, ""))
}

func handleCommand(msg *tgbotapi.Message) {
	switch msg.Command() {
	case "start":
		sendMainMenu(msg.Chat.ID, "Добро пожаловать! Для доступа к функциям войдите в систему.")
	case "help":
		sendMessage(msg.Chat.ID, "Доступные команды:\n/start - начать работу\n/help - помощь")
	}
}

func handleUnauthorized(msg *tgbotapi.Message) {
	switch msg.Text {
	case "🔑 Войти":
		requestLogin(msg.Chat.ID)
	default:
		sendMainMenu(msg.Chat.ID, "Пожалуйста, сначала войдите в систему.")
	}
}

func handleAuthorized(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	user := authorizedUsers[chatID]

	switch msg.Text {
	case "ℹ️ Обо мне":
		sendAboutMeMenu(chatID, user)
	case "🏢 О компании":
		sendMessage(chatID, "Компания: "+user.Company+"\nКонтактные данные: ...")
	case "📂 Проекты":
		projects := "Ваши проекты:\n" + strings.Join(user.Projects, "\n")
		sendMessage(chatID, projects)
	case "🚪 Выйти":
		usersMutex.Lock()
		delete(authorizedUsers, chatID)
		usersMutex.Unlock()
		sendMainMenu(chatID, "Вы вышли из системы.")
	default:
		sendAuthMenu(chatID, "Выберите действие:")
	}
}

func sendAboutMeMenu(chatID int64, user *User) {
	info := fmt.Sprintf("Информация о вас:\nИмя: %s\nЛогин: %s\nКомпания: %s",
		user.Name, user.Username, user.Company)

	// Создаем inline-клавиатуру
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
	bot.Send(msg)
}

func requestLogin(chatID int64) {
	loginState[chatID] = true
	msg := tgbotapi.NewMessage(chatID, "Введите логин и пароль в формате: логин:пароль\nПример: admin:admin123")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}

func processLogin(chatID int64, text string) {
	credentials := strings.SplitN(text, ":", 2)
	if len(credentials) != 2 {
		sendMessage(chatID, "Неверный формат. Используйте: логин:пароль")
		return
	}

	username := strings.TrimSpace(credentials[0])
	password := strings.TrimSpace(credentials[1])

	for _, user := range predefinedUsers {
		if user.Username == username && user.Password == password {
			usersMutex.Lock()
			authorizedUsers[chatID] = &User{
				Username: user.Username,
				Name:     user.Name,
				Company:  user.Company,
				Projects: user.Projects,
			}
			usersMutex.Unlock()
			sendAuthMenu(chatID, "Авторизация успешна! Добро пожаловать, "+user.Name+"!")
			return
		}
	}

	sendMessage(chatID, "Неверный логин или пароль")
	sendMainMenu(chatID, "Попробуйте еще раз или обратитесь к администратору.")
}

func sendMainMenu(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = mainMenuKeyboard
	bot.Send(msg)
}

func sendAuthMenu(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = authMenuKeyboard
	bot.Send(msg)
}

func sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}
