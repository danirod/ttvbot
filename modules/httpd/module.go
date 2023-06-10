package httpd

import (
	"context"
	"net"

	"danirod.es/pkg/ttvbot"
	"go.uber.org/fx"
)

var Module = fx.Module("httpd",
	fx.Provide(fx.Annotate(
		func(conf *ttvbot.Config) *HttpD {
			return newHttpd(conf.HttpBind)
		},
		fx.OnStart(func(ctx context.Context, d *HttpD) error {
			ln, err := net.Listen("tcp", d.server.Addr)
			if err != nil {
				return err
			}

			go d.server.Serve(ln)
			return nil
		}),
		fx.OnStop(func(ctx context.Context, d *HttpD) error {
			d.server.Shutdown(ctx)
			return nil
		}),
	)),
)
