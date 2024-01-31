package main

import (
	"database/sql"
	"log/slog"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/gustapinto/go_hex/cmds/rest/route/ping"
	"github.com/gustapinto/go_hex/internal/account"
	"github.com/gustapinto/go_hex/pkg/httputil"
)

const (
	ServerAddress = "0.0.0.0:8080"
)

func StartServer(db *sql.DB) error {
	mux := http.NewServeMux()

	accountRepository := account.NewSqlDataSource(db)
	accountInteractor := account.NewInteractor(accountRepository)
	_ = accountInteractor

	mux.HandleFunc("GET /v1/ping", httputil.Log(ping.Pong))

	slog.Info("Starting HTTP server", "address", ServerAddress)

	return http.ListenAndServe(ServerAddress, mux)
}
