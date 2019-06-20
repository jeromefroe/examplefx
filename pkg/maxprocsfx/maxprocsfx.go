package maxprocsfx

import (
	"context"

	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module adjusts runtime.GOMAXPROCS to match the CPU quota configured in
// Linux containers. In non-Linux environments, it's a no-op.
//
// In Linux environments without a cgroup configured and without an
// explicitly-set GOMAXPROCS, Module will prevent applications from starting.
// To avoid this, set the GOMAXPROCS environment variable to the desired value
// manually.
var Module = fx.Invoke(Set)

// Params defines the dependencies of the maxprocsfx module.
type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger *zap.Logger
}

// Set uses the provided dependencies to alter runtime concurrency on
// application startup and undo any alterations before shutting down.
func Set(p Params) error {
	undo := func() {}
	p.Lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			var err error
			undo, err = maxprocs.Set(maxprocs.Logger(p.Logger.Sugar().Infof))
			return err
		},
		OnStop: func(context.Context) error {
			undo()
			return nil
		},
	})
	return nil
}
