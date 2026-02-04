package pkg

import (
	"context"
	"fmt"
	"io"
	"os"

	db "github.com/MediStatTech/biometric-service/internal/infra/db"
	pkg_commitplan "github.com/MediStatTech/biometric-service/pkg/commitplan"
	"github.com/MediStatTech/biometric-service/pkg/config"
	"github.com/MediStatTech/commitplan"
	"github.com/MediStatTech/logger"
)

type Facade struct {
	Committer *commitplan.Facade
	Postgres  *db.DB
	Logger    *logger.Logger
	Config    *config.Config
}

func New(ctx context.Context) (*Facade, error) {
	config, err := config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create config: %w", err)
	}

	db, err := initDB(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}

	committer := initCommitter(db)
	logger := initLogger()

	return &Facade{
		Committer: committer,
		Postgres:  db,
		Logger:    logger,
		Config:    config,
	}, nil
}

func initDB(ctx context.Context, config *config.Config) (*db.DB, error) {

	db, err := db.New(&db.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		Database: config.DBName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}

	return db, nil
}

func initCommitter(db *db.DB) *commitplan.Facade {
	return pkg_commitplan.NewCommitter(db.DB)
}

func initLogger() *logger.Logger {
	// TODO: Implement logger configuration
	return logger.New(io.Writer(os.Stdout))
}
