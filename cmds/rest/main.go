package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"github.com/gustapinto/go_hex/cmds/rest/handler"
	"github.com/gustapinto/go_hex/pkg/rest"
)

const (
	ServerAddress = "0.0.0.0:8080"
)

func main() {
	mux := http.NewServeMux()

	ping := handler.Ping{}
	{
		mux.HandleFunc("GET /v1/ping", LogRequest(ping.Pong))
	}

	slog.Info("Starting HTTP server", "address", ServerAddress)

	err := http.ListenAndServe(ServerAddress, mux)
	if err != nil {
		slog.Error("Failed to start http server", "error", err.Error())
		os.Exit(1)
	}
}
