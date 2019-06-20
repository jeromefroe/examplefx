package httpfx

import (
	"context"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module starts an HTTP server and returns a http.ServeMux to use to register handlers
// for the server.
var Module = fx.Provide(New)

// Params defines the dependencies of the httpfx module.
type Params struct {
	fx.In

	Lifecycle  fx.Lifecycle
	Shutdowner fx.Shutdowner
	Logger *zap.Logger
}

// Result defines the objects that the httpfx module provides.
type Result struct {
	fx.Out

	Mux *http.ServeMux
}

// New exports functionality similar to Module, but allows the caller to wrap
// or modify Result. Most users should use Module instead.
func New(p Params) (Result, error) {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "127.0.0.1:8392",
		Handler: mux,
		ErrorLog: zap.NewStdLog(p.Logger),
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := server.ListenAndServe()
				if err != nil && err != http.ErrServerClosed {
					p.Logger.Error("failed to shut down http server cleanly", zap.Error(err))
					p.Shutdowner.Shutdown()
				}
			}()
			p.Logger.Info("starting HTTP server", zap.String("addr", server.Addr))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := server.Shutdown(ctx)
			if err == http.ErrServerClosed {
				return nil
			}
			return err
		},
	})

	return Result{Mux: mux}, nil
}
