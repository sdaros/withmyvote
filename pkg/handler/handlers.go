package withmyvote

import (
	"io/ioutil"
	"net/http"
)

func homePageHandler(rw http.ResponseWriter, req *http.Request) {
	index, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(rw, "Error: Could not open index.html", http.StatusInternalServerError)
		return
	}
	rw.Write(index)
	return
}
