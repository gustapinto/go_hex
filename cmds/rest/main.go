package main

import (
	"log/slog"
	"os"
)

func main() {
	err := StartServer()
	if err != nil {
		slog.Error("Failed to start http server", "error", err.Error())
		os.Exit(1)
	}
}
