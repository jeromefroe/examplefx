package main

import (
	"github.com/jeromefroe/examplefx/pkg/configfx"
	"github.com/jeromefroe/examplefx/pkg/zapfx"
	"github.com/jeromefroe/examplefx/pkg/httpfx"
	"github.com/jeromefroe/examplefx/pkg/maxprocsfx"
	"github.com/jeromefroe/examplefx/pkg/prometheusfx"
	"go.uber.org/fx"
)

// TODO:
//   - rename httpfx to systemfx?
//   - pproffx
//   - versionfx package which reports version as Prometheus metrics
//   - jaegerfx
//   - healthfx
//   - debugfx for debug pages

func main() {
	fx.New(
		configfx.Module,
		httpfx.Module,
		maxprocsfx.Module,
		prometheusfx.Module,
		zapfx.Module,
	).Run()
}
