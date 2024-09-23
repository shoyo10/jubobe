package pgmigration

import (
	"jubobe/internal/model"
	"jubobe/pkg/config"
	"jubobe/pkg/postgres"
	"jubobe/pkg/zerolog"

	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var ServerCmd = &cobra.Command{
	Run: run,
	Use: "pgmigration",
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Error().Msg("need argument: migrate or rollback")
		return
	}

	cfg, err := config.New()
	if err != nil {
		os.Exit(1)
	}

	zerolog.Init(cfg.Log)

	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Error().Msgf("Error connecting to database, %s", err)
		os.Exit(1)
	}

	if err := migration(db, args[0]); err != nil {
		log.Error().Msgf("Error migrating, %s", err)
		os.Exit(1)
	}
}

func migration(db *gorm.DB, cmd string) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{{
		ID: "202409240225",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.Migrator().CreateTable(&model.Patient{}, &model.Order{}); err != nil {
				return err
			}
			initialPatients := []model.Patient{
				{Name: "Alice"}, {Name: "Bob"}, {Name: "Charlie"}, {Name: "David"}, {Name: "Eve"},
			}
			if err := tx.Create(&initialPatients).Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable((&model.Order{}).TableName(), (&model.Patient{}).TableName())
		},
	}})

	switch cmd {
	case "migrate":
		return m.Migrate()
	case "rollback":
		return m.RollbackLast()
	default:
		log.Error().Msgf("Invalid command %s", cmd)
	}
	return nil
}
