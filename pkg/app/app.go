package app

import (
	"net/http"

	"github.com/sdaros/withmyvote/pkg/handler"
)

var (
	// Routes to handle
	Routes = []Route{
		Route{"/", handler.Root{}},
	}
)

// Route is cool
type Route struct {
	Path    string
	Handler http.Handler
}
