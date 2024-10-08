package apiserver

import (
	"context"
	"jubobe/internal/delivery/http"
	"jubobe/internal/repository/pg"
	"jubobe/internal/service"
	"jubobe/pkg/config"
	"jubobe/pkg/echorouter"
	"jubobe/pkg/postgres"
	"jubobe/pkg/zerolog"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// ServerCmd ...
var ServerCmd = &cobra.Command{
	Run: run,
	Use: "jubobe",
}

func run(cmd *cobra.Command, args []string) {
	cfg, err := config.New()
	if err != nil {
		os.Exit(1)
	}

	zerolog.Init(cfg.Log)

	app := fx.New(
		fx.Supply(*cfg),
		fx.Provide(
			echorouter.FxNewEcho,
			postgres.New,
			pg.New,
			service.New,
			http.NewHandler,
		),
		fx.Invoke(
			http.SetRoutes,
		),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Error().Msg(err.Error())
		return
	}

	<-app.Done()

	log.Info().Msgf("main: shutting down %s...", cmd.Name())

	stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.Stop(stopCtx); err != nil {
		log.Error().Msg(err.Error())
	}
}
