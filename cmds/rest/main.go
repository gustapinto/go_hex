package main

import (
	"github.com/gustapinto/go_hex/cmds/rest/route/ping"
	"github.com/gustapinto/go_hex/pkg/httputil"
	"log/slog"
	"net/http"
	"os"
)

const (
	ServerAddress = "0.0.0.0:8080"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/ping", httputil.Log(ping.Pong))

	slog.Info("Starting HTTP server", "address", ServerAddress)

	err := http.ListenAndServe(ServerAddress, mux)
	if err != nil {
		slog.Error("Failed to start http server", "error", err.Error())
		os.Exit(1)
	}
}
