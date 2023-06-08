package database

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Module("database",
	fx.Provide(fx.Annotate(
		newDatabase,
		fx.OnStart(func(ctx context.Context, db *Database) error {
			return db.init()
		}),
		fx.OnStop(func(ctx context.Context, db *Database) error {
			return db.close()
		}),
	)),
)
