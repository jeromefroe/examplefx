package jaegerfx

import (
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jconfig "github.com/uber/jaeger-client-go/config"
	"go.uber.org/config"
	"go.uber.org/fx"
)

const (
	// ConfigurationKey is the portion of the service configuration that this
	// package reads.
	ConfigurationKey = "tracing"
)

// Module provides an opentracing.Tracer, and it also configures opentracing's
// package-global state. It reads its configuration from the "tracing" key of the
// service configuration.
//
// In YAML, tracing configuration might look like this:
//
//  tracing:
//    TODO
var Module = fx.Options(
// fx.Provide(New, newHTTPMiddleware),
// fx.Invoke(setGlobalTracer),
)

// Params defines the dependencies of the jaegerfx module.
type Params struct {
	fx.In

	Config config.Provider
}

// Result defines the objects that the jaegerfx module provides.
type Result struct {
	fx.Out

	Tracer opentracing.Tracer
}

// New exports functionality similar to Module, but allows the caller to wrap
// or modify Result. Most users should use Module instead.
func New(p Params) (Result, error) {
	// jaegerConfig := ...

	return Result{}, nil
}
