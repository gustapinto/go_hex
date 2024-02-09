package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/gustapinto/go_hex/cmds/rest/handler"
	"github.com/gustapinto/go_hex/internal/datasource/database"
	"github.com/gustapinto/go_hex/internal/interactor"
	"github.com/gustapinto/go_hex/pkg/httputil"
	_ "modernc.org/sqlite"
)

const (
	DatabaseDriver = "sqlite"
	DatabaseDSN    = "./database.db"
	ServerAddress  = "0.0.0.0:8080"
)

func main() {
	db, err := sql.Open(DatabaseDriver, DatabaseDSN)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err.Error())
		os.Exit(1)
	}

	mux := http.NewServeMux()

	accountHandler := handler.NewAccount(interactor.NewAccount(database.NewAccount(db)))
	{
		mux.HandleFunc("GET /v1/accounts", httputil.Log(accountHandler.Get))
		mux.HandleFunc("GET /v1/accounts/{id}", httputil.Log(accountHandler.GetByID))
		mux.HandleFunc("POST /v1/accounts", httputil.Log(accountHandler.Create))
		mux.HandleFunc("PUT /v1/accounts/{id}", httputil.Log(accountHandler.UpdateByID))
		mux.HandleFunc("DELETE /v1/accounts/{id}", httputil.Log(accountHandler.DeletebyID))
	}

	slog.Info("Starting HTTP server", "address", ServerAddress)

	err = http.ListenAndServe(ServerAddress, mux)
	if err != nil {
		slog.Error("Failed to start http server", "error", err.Error())
		os.Exit(1)
	}
}
