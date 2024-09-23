package postgres

import (
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v3"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config ...
type Config struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Username string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	DBName   string `yaml:"dbname" mapstructure:"dbname"`
}

func New(cfg *Config) (*gorm.DB, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Taipei",
		cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port)
	dialector := postgres.Open(dsn)

	var conn *gorm.DB
	err := backoff.Retry(func() error {
		db, err := gorm.Open(dialector, &gorm.Config{
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
		})
		if err != nil {
			return err
		}
		conn = db

		sqlDB, err := conn.DB()
		if err != nil {
			return err
		}

		err = sqlDB.Ping()
		return err
	}, bo)

	if err != nil {
		log.Error().Msgf("main: database connect err: %s", err.Error())
		return nil, err
	}
	log.Info().Msgf("database ping success")

	sqlDB, err := conn.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(600 * time.Second)

	return conn, nil
}
