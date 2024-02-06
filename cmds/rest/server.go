package main

import (
	"database/sql"
	"log/slog"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/gustapinto/go_hex/cmds/rest/handler"
	"github.com/gustapinto/go_hex/internal/datasource/database"
	"github.com/gustapinto/go_hex/internal/interactor"
	"github.com/gustapinto/go_hex/pkg/httputil"
)

const (
	ServerAddress = "0.0.0.0:8080"
)

func StartServer(db *sql.DB) error {
	mux := http.NewServeMux()

	accountHandler := handler.NewAccount(interactor.NewAccount(database.NewAccount(db)))
	{
		mux.HandleFunc("GET /v1/account", httputil.Log(accountHandler.Get))
		mux.HandleFunc("POST /v1/account", httputil.Log(accountHandler.Create))
		mux.HandleFunc("GET /v1/account/{id}", httputil.Log(accountHandler.GetByID))
		mux.HandleFunc("PUT /v1/account/{id}", httputil.Log(accountHandler.UpdateByID))
		mux.HandleFunc("DELETE /v1/account/{id}", httputil.Log(accountHandler.DeletebyID))
	}

	slog.Info("Starting HTTP server", "address", ServerAddress)

	return http.ListenAndServe(ServerAddress, mux)
}
