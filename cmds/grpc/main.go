package main

import (
	"database/sql"
	"log/slog"
	"net"
	"os"

	"github.com/gustapinto/go_hex/cmds/grpc/gen"
	"github.com/gustapinto/go_hex/cmds/grpc/server"
	"github.com/gustapinto/go_hex/internal/datasource/database"
	"github.com/gustapinto/go_hex/internal/interactor"
	"google.golang.org/grpc"
)

const (
	DatabaseDriver = "sqlite"
	DatabaseDSN    = "./database.db"
	ServerAddress  = "0.0.0.0:8081"
)

func main() {
	db, err := sql.Open(DatabaseDriver, DatabaseDSN)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err.Error())
		os.Exit(1)
	}

	migration := database.NewMigration(db)
	if err := migration.Up(); err != nil {
		slog.Error("Failed to run migration", "error", err.Error())
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		slog.Error("Failed to open listener", "error", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()

	accountRepository := database.NewAccount(db)
	transactionRepository := database.NewTransaction(db)
	accountInteractor := interactor.NewAccount(accountRepository, transactionRepository)
	accountServer := server.NewAccount(accountInteractor)
	gen.RegisterAccountServiceServer(grpcServer, &accountServer)

	if err := grpcServer.Serve(listener); err != nil {
		slog.Error("Failed to attach gRPC server", "error", err.Error())
		os.Exit(1)
	}
}
