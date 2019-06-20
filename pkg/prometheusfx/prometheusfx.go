package prometheusfx

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	// ConfigurationKey is the portion of the configuration that this package reads.
	ConfigurationKey = "prometheus"

	_defaultEndpoint = "/metrics"
)

// Module registers a Prometheus metrics handler.
//
// In YAML, the prometheus configuration might look like this:
//
// prometheus:
//   endpoint: "/foo"
var Module = fx.Invoke(Register)

// Configuration controls option for the Prometheus handler. All fields are optional.
type Configuration struct {
	Endpoint string `yaml:"endpoint"` // Defaults to "/metrics".
}

func newConfiguration(cfg config.Provider) (Configuration, error) {
	var c Configuration

	if err := cfg.Get(ConfigurationKey).Populate(&c); err != nil {
		return Configuration{}, fmt.Errorf("failed to set up configuration: %v", err)
	}

	if c.Endpoint == "" {
		c.Endpoint = _defaultEndpoint
	}

	return c, nil
}

// Params defines the dependencies of the prometheusfx module.
type Params struct {
	fx.In

	Config config.Provider
	Logger *zap.Logger
	Mux    *http.ServeMux
}

// Register adds a Prometheus metrics handler to the *http.ServeMux provided in p.
func Register(p Params) error {
	c, err := newConfiguration(p.Config)
	if err != nil {
		return fmt.Errorf("failed to load prometheus configuration: %v", err)
	}

	p.Mux.Handle(c.Endpoint, promhttp.Handler())
	p.Logger.Info("prometheus: registered handler", zap.String("endpoint", c.Endpoint))

	return nil
}
