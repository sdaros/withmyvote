package main

import (
	"crypto/tls"
	"net/http"
	"sync"
)

type DB interface {
	LoadOrStore(key string, value interface{})
}

var store = struct {
	state sync.Map
}{
	state: sync.Map{},
}

func (db *store) LoadOrStore(key string, value interface{}) {
	db.state.LoadOrStore(key, value)
}

func init() {
	/*
	   # parse flags

	   - flag: server.crt and server.key
	*/
}

func main() {
	server := HTTPServerWithSecureTLSOptions()
	server.Addr, _ = db.LoadOrStore("server.port", ":8080")
	mux := http.NewServeMux()
	/*
		mux.HandleFunc("/", homePageHandler)
		mux.HandleFunc("/choose/0", deleteVoteHandler)
		mux.HandleFunc("/choose/", storeVoteHandler)
	*/
	server.Handler = mux
}

// HTTPServerWithSecureTLSOptions returns a http server configured to use
// secure cipher suites and curves as defined by the german federal office
// for information security (BSI) in TR-02102-2 version 2018-01
func HTTPServerWithSecureTLSOptions() *http.Server {
	cfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.CurveP521,
			tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		},
	}
	return &http.Server{
		TLSConfig: cfg,
		TLSNextProto: make(map[string]func(*http.Server,
			*tls.Conn, http.Handler), 0),
	}
}
