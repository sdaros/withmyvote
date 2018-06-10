package handler

import (
	"net/http"
)

// Root handler acts as the default fallthrough handler
type Ballot struct{}

func (b *Ballot) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Cast your ballot here"))
}
