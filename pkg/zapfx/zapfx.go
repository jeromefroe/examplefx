package zapfx

import (
	"context"
	"fmt"

	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	// ConfigurationKey is the portion of the configuration that this package reads.
	ConfigurationKey = "logging"
)

// Module provides a zap logger for structured logging.
//
// In YAML, logging configuration might look like this:
//
//  logging:
//    level: info
//    development: false
//    sampling:
//      initial: 100
//      thereafter: 100
//    encoding: json
var Module = fx.Provide(New)

// Params defines the dependencies of the zapfx module.
type Params struct {
	fx.In

	Config    config.Provider
	Lifecycle fx.Lifecycle
}

// Result defines the objects that the zapfx module provides.
type Result struct {
	fx.Out

	Level  zap.AtomicLevel
	Logger *zap.Logger
}

// New exports functionality similar to Module, but allows the caller to wrap
// or modify Result. Most users should use Module instead.
func New(p Params) (Result, error) {
	var (
		c   = zap.NewProductionConfig()
		raw = p.Config.Get(ConfigurationKey)
	)
	if err := raw.Populate(&c); err != nil {
		return Result{}, fmt.Errorf("failed to load logging config: %v", err)
	}

	logger, err := c.Build()
	if err != nil {
		return Result{}, fmt.Errorf("failed to create zap logger: %v", err)
	}

	p.Lifecycle.Append(fx.Hook{
		OnStop: func(context.Context) error {
			logger.Sync()
			return nil
		},
	})

	return Result{
		Level:  c.Level,
		Logger: logger,
	}, nil
}
