package main

import (
	"flag"
	"log"

	"github.com/sdaros/withmyvote/pkg/server"
)

var (
	certFile string
	keyFile  string
	port     string
)

func init() {
	flag.StringVar(&certFile, "cert-file", "server.crt", "location of public tls certificate file")
	flag.StringVar(&keyFile, "key-file", "server.key", "location of private tls certificate file")
	flag.StringVar(&port, "port", ":8080", "port number to listen on")
	flag.Parse()
}

func main() {
	svr := server.New(port)
	server.Handler = server.Register(routes...)
	log.Fatal(svr.ListenAndServe())
}
