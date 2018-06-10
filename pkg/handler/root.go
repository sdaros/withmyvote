package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Root handler acts as the default fallthrough handler
type Root struct{}

func (r Root) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if len(strings.Split(req.URL.Path, "/")[1]) > 3 {
		b := &Ballot{}
		b.ServeHTTP(rw, req)
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	http.FileServer(
		http.Dir(fmt.Sprintf("%s/web/static", wd))).
		ServeHTTP(rw, req)
	return
}
