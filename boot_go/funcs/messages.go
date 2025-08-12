package funcs

import (
	"fmt"
	"time"

	provider "../providers"
)

func sendTelegramMessage(text string, chatID int64) error {

	SendAuthMenu(chatID, "message")

	return nil
}

func startTelegramBot(interval time.Duration, message string, chatID int64) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			_, exists := provider.AuthorizedUsers[chatID]
			if !exists {
				return
			}
			err := sendTelegramMessage(message, chatID)
			if err != nil {
				fmt.Printf("Ошибка отправки: %v\n", err)
			}

		}
	}

}
