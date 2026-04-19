package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/twenty99h/wishes-service/config"
	"github.com/twenty99h/wishes-service/internal/controller/http"
	"github.com/twenty99h/wishes-service/internal/usecase"
	"github.com/twenty99h/wishes-service/pkg/httpserver"
	"github.com/twenty99h/wishes-service/pkg/postgres"
	"github.com/twenty99h/wishes-service/pkg/router"
)

func Run(ctx context.Context, config config.Config) error {
	pgPool, err := postgres.New(ctx, config.Postgres)
	if err != nil {
		return fmt.Errorf("pgPool.New: %w", err)
	}

	uc := usecase.New(pgPool)

	r := router.New()
	http.WishesRouter(r, uc)
	httpServer := httpserver.New(r, config.HTTP)

	log.Info().Msg("app started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	log.Info().Msg("app got signal to shutdown")
	httpServer.Close()
	log.Info().Msg("app shutting down")

	return nil
}
