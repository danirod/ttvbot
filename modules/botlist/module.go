package botlist

import (
	"context"

	"danirod.es/pkg/ttvbot"
	"go.uber.org/fx"
)

var Module = fx.Module("botlist",
	fx.Provide(fx.Annotate(
		func(conf *ttvbot.Config) *Botlist {
			list := newBotList()
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
