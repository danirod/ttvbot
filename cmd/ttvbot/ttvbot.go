package main

import (
	"danirod.es/pkg/ttvbot"
	"danirod.es/pkg/ttvbot/invokes/botdetector"
	"danirod.es/pkg/ttvbot/invokes/reply"
	"danirod.es/pkg/ttvbot/invokes/roster"
	"danirod.es/pkg/ttvbot/modules/botlist"
	"danirod.es/pkg/ttvbot/modules/chatbot"
	"danirod.es/pkg/ttvbot/modules/database"
	"danirod.es/pkg/ttvbot/modules/httpd"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	app := fx.New(
		fx.Provide(ttvbot.LoadConfig),
		database.Module,
		httpd.Module,
		botlist.Module,
		chatbot.Module,
		fx.Invoke(roster.Invoke),
		fx.Invoke(botdetector.Invoke),
		fx.Invoke(reply.Invoke),
	)
	app.Run()
}
