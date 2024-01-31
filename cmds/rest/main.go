package main

import (
	"log/slog"
	"os"
)

func main() {
	db, err := StartDatabase()
	if err != nil {
		slog.Error("Failed to connect to database", "error", err.Error())
		os.Exit(1)
	}

	err = StartServer(db)
	if err != nil {
		slog.Error("Failed to start http server", "error", err.Error())
		os.Exit(1)
	}
}
