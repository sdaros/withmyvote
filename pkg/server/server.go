package withmyvote

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/sdaros/withmyvote/pkg/app"
)

func New(port string) *http.Server {
	server := HTTPServerWithSecureTLSOptions()
	if port != "" {
		server.Addr = port
	}
	return server
}
func Register(routes ...app.Route) http.Handler {
	mux := http.NewServeMux()
	for _, r := range routes {
		mux.Handle(r.path, r.handler)
	}
	return mux
}

// HTTPServerWithSecureTLSOptions returns a http server configured to use
// secure cipher suites and curves as defined by the german federal office
// for information security (BSI) in TR-02102-2 version 2018-01
func HTTPServerWithSecureTLSOptions() *http.Server {
	cfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    cfg,
	}
}
