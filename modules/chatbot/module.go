package chatbot

import (
	"context"

	"danirod.es/pkg/ttvbot"
	"go.uber.org/fx"
)

var Module = fx.Module("chatbot",
	fx.Provide(fx.Annotate(
		func(conf *ttvbot.Config) *Chatbot {
			return newChatbot(conf.BotUsername, conf.IrcToken, conf.TargetChannel)
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
