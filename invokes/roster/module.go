package roster

import (
	"danirod.es/pkg/ttvbot/modules/chatbot"
	"danirod.es/pkg/ttvbot/modules/database"
)

func Invoke(db *database.Database, bot *chatbot.Chatbot) {
	bot.OnUserJoin(func(channel, username string) {
		if err := db.JoinPartDB.Join(username, channel); err != nil {
			panic(err)
		}
	})
	bot.OnUserLeave(func(channel, username string) {
		if err := db.JoinPartDB.Part(username, channel); err != nil {
			panic(err)
		}
	})
	bot.OnMessage(func(channel, username, content string, flags *chatbot.MessageFlags) {
		if err := db.MessageDB.Message(username, channel, content, flags.First); err != nil {
			panic(err)
		}
	})
}
