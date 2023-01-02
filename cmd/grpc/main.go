package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/suryaadi44/iris-playground/app/api/grpc/bootstrapper"
	"github.com/suryaadi44/iris-playground/app/api/grpc/pb"
	"github.com/suryaadi44/iris-playground/utils/config"
	"github.com/suryaadi44/iris-playground/utils/database/postgresql"
	"github.com/suryaadi44/iris-playground/utils/shutdown"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.Load("config")
	if err != nil {
		log.Fatalf("[Config] Error loading config: %v", err)
	}

	local, err := time.LoadLocation(config.GetString("timezone"))
	if err != nil {
		log.Fatalf("[Config] Error loading timezone: %v", err)
	}
	time.Local = local

	db := postgresql.InitDatabase(
		config.GetString("database.postgres.host"),
		config.GetInt("database.postgres.port"),
		config.GetString("database.postgres.user"),
		config.GetString("database.postgres.password"),
		config.GetString("database.postgres.database"),
		config.GetString("timezone"),
	)

	authService := bootstrapper.InitGRPC(db, config)

	listen := net.JoinHostPort(config.GetString("grpc.host"), config.GetString("grpc.port"))
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("[Server] Error running server: %v", err)
	}
	log.Printf("[Server] Running on %s", listen)

	g := grpc.NewServer()
	pb.RegisterAuthenticateServer(g, authService)
	go func() {
		if err := g.Serve(listener); err != nil {
			log.Fatalf("[Server] Error running server: %v", err)
		}
	}()

	ctx := context.Background()
	// operations to be executed on shutdown
	ops := map[string]shutdown.Operation{
		// TODO: Add more operations here
		// "redis": ...

		"database": func(ctx context.Context) error {
			dbCon, err := db.DB()
			if err != nil {
				return err
			}
			return dbCon.Close()
		},
		"grpc-server": func(ctx context.Context) error {
			g.GracefulStop()
			return nil
		},
	}
	// listen for interrupt signal
	<-shutdown.GracefulShutdown(ctx, config.GetDuration("shutdown_timeout"), ops)
}
