package main

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/twenty99h/wishes-service/config"
	"github.com/twenty99h/wishes-service/internal/app"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}
	ctx := context.Background()

	if err = app.Run(ctx, c); err != nil {
		log.Fatal().Err(err).Msg("app.Run")
	}
}
