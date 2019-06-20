package zapfx

import (
	"fmt"
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	// ConfigurationKey is the portion of the configuration that this package reads.
	ConfigurationKey = "logging"
)

// Module provides a zap logger for structured logging.
// In production and staging, the default configuration logs at zap.InfoLevel,
// uses the JSON encoder, and enables sampling. In all other environments, the
// default configuration logs at zap.DebugLevel, uses the console encoder, and
// disables sampling.
//
// In YAML, logging configuration might look like this:
//
// TODO
var Module = fx.Provide(New)

// Params defines the dependencies of the zapfx module.
type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
}

// Result defines the objects that the zapfx module provides.
type Result struct {
	fx.Out

	// TODO:
	// Level         zap.AtomicLevel
	// Config        zap.Config
	Logger *zap.Logger
}

// New exports functionality similar to Module, but allows the caller to wrap
// or modify Result. Most users should use Module instead.
func New(params Params) (Result, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Result{}, fmt.Errorf("failed to create zap logger: %v", err)
	}

	params.Lifecycle.Append(fx.Hook{
		OnStop: func(context.Context) error {
			logger.Sync()
			return nil
		},
	})

	return Result{Logger: logger}, nil
}
