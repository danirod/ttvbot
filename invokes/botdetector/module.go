package botdetector

import (
	"fmt"
	"math/rand"
	"strings"

	"danirod.es/pkg/ttvbot/modules/botlist"
	"danirod.es/pkg/ttvbot/modules/chatbot"
)

var suspiciousSentences = []string{
	"¿me puedes decir cuánto es 2+2?",
	"¿cuántas motos ves en la siguiente imagen?",
	"¿qué números ves en el siguiente captcha?",
	"¿no serás por casualidad el primo de ChatGPT?",
	"¿me puedes responder a un par de preguntas?",
	"¿tienes sentimientos?",
}

func Invoke(bot *chatbot.Chatbot, list *botlist.Botlist) {
	bot.OnUserJoin(func(channel, username string) {
		fmt.Printf("%s joined %s.\n", username, channel)
		if list.IsBot(username) {
			suspicious := suspiciousSentences[rand.Intn(len(suspiciousSentences))]
			msg := fmt.Sprintf("%s, %s", username, suspicious)
			bot.SendMessage(msg)
		}
	})
}
