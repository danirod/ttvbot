package reply

import (
	"fmt"

	"danirod.es/pkg/ttvbot/modules/chatbot"
)

func Invoke(bot *chatbot.Chatbot) {
	bot.OnMessage(func(channel, username, message string, flags *chatbot.MessageFlags) {
		if message == "!status" {
			msg := fmt.Sprintf("@%s, el bot está vivo", username)
			bot.SendMessage(msg)
		}
	})
}
