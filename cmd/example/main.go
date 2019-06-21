package main

import (
	"github.com/jeromefroe/examplefx/pkg/configfx"
	"github.com/jeromefroe/examplefx/pkg/httpfx"
	"github.com/jeromefroe/examplefx/pkg/maxprocsfx"
	"github.com/jeromefroe/examplefx/pkg/pproffx"
	"github.com/jeromefroe/examplefx/pkg/prometheusfx"
	"github.com/jeromefroe/examplefx/pkg/zapfx"
	"go.uber.org/fx"
)

// TODO:
//   - jaegerfx
//   - healthfx
//     - need something to signal when application is ready to be able to applications that
//       need to warmup, is a shim around a WaitGroup enough?
//     - also probably want some cooldown period
//     - states: is Healthy and Unhealthy enough? Worth distinguishing Starting, Healthy,
//       Unhealthy, Stopping, Stopped?
//   - rename httpfx to systemfx? used to register administrative and introspection handlers
//   - metadatafx with build info? e.g. hash, time, user, host
//   - debugfx might be nice so users can add consistent pages, need a Route type. e.g. a
//     debug route for the fx dependencies (dot module) is nifty

func main() {
	fx.New(
		configfx.Module,
		httpfx.Module,
		maxprocsfx.Module,
		pproffx.Module,
		prometheusfx.Module,
		zapfx.Module,
	).Run()
}
