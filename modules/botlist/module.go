package botlist

import (
	"context"
	"fmt"

	"danirod.es/pkg/ttvbot"
	"danirod.es/pkg/ttvbot/modules/httpd"
	"go.uber.org/fx"
)

var Module = fx.Module("botlist",
	fx.Provide(fx.Annotate(
		func(conf *ttvbot.Config, http *httpd.HttpD) *Botlist {
			list := newBotList()
			http.AddHealthCheck("botlist", func() (bool, string) {
				bots := fmt.Sprintf("%d bots tracked", list.CountBots())
				return true, bots
			})
			list.AllowBot(conf.AllowedBots...)
			return list
		},
		fx.OnStart(func(ctx context.Context, list *Botlist) error {
			return list.tryRefresh()
		}),
		fx.OnStop(func(ctx context.Context) error {
			return nil
		}),
	)),
)
