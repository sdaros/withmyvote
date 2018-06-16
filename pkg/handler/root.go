package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	base = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="css/site.css" type="text/css" media="screen" />
  <title>With My Vote!</title>
</head>
<body>
%s
</body>
</html>
`
)

// Root handler acts as the default fallthrough handler
type Root struct{}

func (r Root) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println("begin")
	if len(strings.Split(req.URL.Path, "/")[1]) > 3 {
		log.Println("> 3")
		b := &Ballot{}
		b.ServeHTTP(rw, req)
		return
	}
	wd, err := os.Getwd()
	log.Println(err)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(
		http.Dir(fmt.Sprintf("%s/web/static", wd))).
		ServeHTTP(rw, req)
}
