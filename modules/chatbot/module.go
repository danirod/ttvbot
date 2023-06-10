package chatbot

import (
	"context"

	"danirod.es/pkg/ttvbot"
	"danirod.es/pkg/ttvbot/modules/httpd"
	"go.uber.org/fx"
)

var Module = fx.Module("chatbot",
	fx.Provide(fx.Annotate(
		func(conf *ttvbot.Config, http *httpd.HttpD) *Chatbot {
			chatbot := newChatbot(conf.BotUsername, conf.IrcToken, conf.TargetChannel)
			http.AddHealthCheck("chatbot", func() (bool, string) {
				connected := chatbot.connected.Load()
				if connected {
					return true, "Bot is connected"
				}
				return false, "Bot is not connected at the moment"
			})
			return chatbot
		},
		fx.OnStart(func(ctx context.Context, bot *Chatbot) error {
			go bot.connect()
			return nil
		}),
		fx.OnStop(func(ctx context.Context, bot *Chatbot) error {
			return bot.disconnect()
		}),
	)),
)
