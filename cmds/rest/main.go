package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/gustapinto/go_hex/cmds/rest/handler"
	"github.com/gustapinto/go_hex/cmds/rest/middleware"
	"github.com/gustapinto/go_hex/internal/datasource/database"
	"github.com/gustapinto/go_hex/internal/interactor"
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

	accountRepository := database.NewAccount(db)
	transactionRepository := database.NewTransaction(db)
	accountInteractor := interactor.NewAccount(accountRepository, transactionRepository)
	accountHandler := handler.NewAccount(accountInteractor)
	{
		mux.HandleFunc("GET /v1/accounts", accountHandler.Get)
		mux.HandleFunc("GET /v1/accounts/{accountID}", accountHandler.GetByID)
		mux.HandleFunc("POST /v1/accounts", accountHandler.Create)
		mux.HandleFunc("PUT /v1/accounts/{accountID}", accountHandler.UpdateByID)
		mux.HandleFunc("DELETE /v1/accounts/{accountID}", accountHandler.DeletebyID)
		mux.HandleFunc("GET /v1/accounts/{accountID}/transactions", accountHandler.GetTransactionsByAccountID)
		mux.HandleFunc("GET /v1/accounts/{accountID}/transactions/{transactionID}", accountHandler.GetTransactionByIDAndAccountID)
		mux.HandleFunc("POST /v1/accounts/{accountID}/transactions", accountHandler.CreateTransaction)
		mux.HandleFunc("DELETE /v1/accounts/{accountID}/transactions/{transactionID}", accountHandler.DeleteTransactionByIDAndAccountID)
	}

	slog.Info("Starting HTTP server", "address", ServerAddress)

	muxWithLogger := middleware.WrapWithLogger(mux)

	err = http.ListenAndServe(ServerAddress, muxWithLogger)
	if err != nil {
		slog.Error("Failed to start http server", "error", err.Error())
		os.Exit(1)
	}
}
