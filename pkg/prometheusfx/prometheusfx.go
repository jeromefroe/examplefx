package prometheusfx

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	// ConfigurationKey is the portion of the configuration that this package reads.
	ConfigurationKey = "metrics"

	_defaultEndpoint = "/metrics"
)

// Module registers a Prometheus metrics handler on the provided ServeMux.
//
// In YAML, the metrics configuration might look like this:
//
// metrics:
//   endpoint: "/custom_path"
var Module = fx.Invoke(registerHandler)

// Configuration controls option for the Prometheus handler. All fields are optional.
type Configuration struct {
	Endpoint      string `yaml:"endpoint"` // Defaults to "/metrics".
	DropBuildInfo bool   `yaml:"dropBuildInfo"`
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
	Mux    *http.ServeMux

	Logger *zap.Logger `optional:"true"`
}

func registerHandler(p Params) error {
	c, err := newConfiguration(p.Config)
	if err != nil {
		return fmt.Errorf("failed to load prometheus configuration: %v", err)
	}

	if p.Logger == nil {
		p.Logger = zap.NewNop()
	}
	p.Logger = p.Logger.With(zap.Namespace("prometheusfx"))

	if !c.DropBuildInfo {
		if err := prometheus.Register(prometheus.NewBuildInfoCollector()); err != nil {
			return fmt.Errorf("failed to register Prometheus build information collector: %v", err)
		}
	}

	p.Mux.Handle(c.Endpoint, promhttp.Handler())
	p.Logger.Info("registered metrics handler", zap.String("endpoint", c.Endpoint))

	return nil
}
