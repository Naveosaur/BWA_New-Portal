package config

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*Postgres, error) {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.PsqlDB.User, cfg.PsqlDB.Password, cfg.PsqlDB.Host, cfg.PsqlDB.Port, cfg.PsqlDB.DBName)
	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
        log.Error().Err(err).Msg("Failed to connect to database")
		return nil, err
	}

    sqlDb, err := db.DB()
    if err != nil {
        log.Error().Err(err).Msg("Failed to get database")
        return nil, err
    }

    sqlDb.SetMaxIdleConns(cfg.PsqlDB.DBMaxIdleConns)
    sqlDb.SetMaxOpenConns(cfg.PsqlDB.DBMaxOpenConns)
    sqlDb.SetConnMaxLifetime(time.Hour)
    
	return &Postgres{DB: db}, nil
}
