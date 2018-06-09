package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	// NEXT serve static content from web/static dir
	wd = func() string {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		return wd
	}()
	// Routes to handle
	Routes = []Route{
		Route{"/", http.FileServer(http.Dir(fmt.Sprintf("%s/web/static", wd)))},
	}
)

// Route is cool
type Route struct {
	Path    string
	Handler http.Handler
}
