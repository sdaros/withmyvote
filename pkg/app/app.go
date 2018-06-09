package app

import "net/http"

var (
	// NEXT serve static content from web/static dir
	routes = []Route{
		Route{"/", http.FileServer(http.Dir("/tmp"))},
	}
)

// Route is cool
type Route struct {
	path    string
	handler http.Handler
}
