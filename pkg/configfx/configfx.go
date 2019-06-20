package configfx

import (
	"go.uber.org/config"
	"go.uber.org/fx"
)

// Module provides a config.Provider.
var Module = fx.Provide(New)

// Params defines the dependencies of the configfx module.
type Params struct {
	fx.In
}

// Result defines the objects that the configfx module provides.
type Result struct {
	fx.Out

	Provider config.Provider
}

// New exports functionality similar to Module, but allows the caller to wrap
// or modify Result. Most users should use Module instead.
func New(p Params) (Result, error) {
	return Result{Provider: config.NopProvider{}}, nil
}
