package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/sdaros/withmyvote/pkg/server"
)

type route struct {
	path    string
	handler http.Handler
}

var (
	certFile string
	keyFile  string
	port     string
	routes   = []route{
		route{"/", http.FileServer(http.Dir("/tmp"))},
	}
)

func init() {
	flag.StringVar(&certFile, "cert-file", "server.crt", "location of public tls certificate file")
	flag.StringVar(&keyFile, "key-file", "server.key", "location of private tls certificate file")
	flag.StringVar(&port, "port", ":8080", "port number to listen on")
	flag.Parse()
}

func main() {
	// NEXT make these changes count
	svr := server.New(port)
	server.Handler = server.Register(routes...)
	log.Fatal(svr.ListenAndServe())
	/*
		server.ListenAndServeTLS(
				db.LoadOrStore("certfile", "./server.crt"),
				db.LoadOrStore("keyfile", "./server.key"),
			)
	*/
}
