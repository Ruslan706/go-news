package main

import (
	"context"
	"fmt"

	"github.com/Ruslan706/go-news/commit/86c72d46a63efd17c0da10c052fdbd455e194421/internal/config"
	"github.com/Ruslan706/go-news/commit/86c72d46a63efd17c0da10c052fdbd455e194421/internal/server"
	"github.com/Ruslan706/go-news/commit/86c72d46a63efd17c0da10c052fdbd455e194421/repository"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.ReadConfig()
	log.Info().Msg("Start service")
	repo := repository.New()
	log.Debug().Any("repo", repo).Msg("Check new repo")
	server := server.New(repo)
	log.Debug().Any("server", server).Msg("Check new server")
	conn, err := initDB(cfg.DbAddr)
	if err != nil {
		panic(err)
	}
	storage := repository.NewDB(conn)
}

func initDB(addr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), addr)
	if err != nil {
		return nil, fmt.Errorf("database intialization error: %w", err)
	}
	return conn, nil
}
