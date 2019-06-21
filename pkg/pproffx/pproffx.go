package pproffx

import (
	"net/http"
	_ "net/http/pprof" // Register pprof handler on DefaultServeMux.

	"go.uber.org/fx"
)

// Module registers pprof handlers on the provided ServeMux.
var Module = fx.Invoke(registerHandler)

// Params defines the dependencies of the pproffx module.
type Params struct {
	fx.In

	Mux *http.ServeMux
}

func registerHandler(p Params) error {
	// Since all the pprof handlers are registered on http.DefaultServeMux
	// as a side effect of importing net/http/pprof, we redirect everything
	// under /debug/pprof to the default serve mux.
	p.Mux.Handle("/debug/pprof/", http.DefaultServeMux)
	return nil
}
