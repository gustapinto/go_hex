package main

import (
	"log/slog"
	"net/http"

	"github.com/gustapinto/go_hex/cmds/rest/route/ping"
	"github.com/gustapinto/go_hex/pkg/httputil"
)

const (
	ServerAddress = "0.0.0.0:8080"
)

func StartServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/ping", httputil.Log(ping.Pong))

	slog.Info("Starting HTTP server", "address", ServerAddress)

	return http.ListenAndServe(ServerAddress, mux)
}
