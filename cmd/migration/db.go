package migration

import (
	"go-shorten/internal/bootstrap"
	"go-shorten/pkg/dbmigration"
)

func MigrateDatabase() {
	cfg := bootstrap.NewConfig()

	dbmigration.DatabaseMigration(cfg)
}
