package migrations

import (
	"Go-Rampup/config"
	"context"
	"database/sql"
	"gorm.io/gorm"
	base_models "Go-Rampup/db/models/base"
	"gorm.io/driver/postgres"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(up0001, down0001)
}

func up0001(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	config := config.GetConfig()
	gormDB, err := gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	gormDB.Migrator().CreateTable(&base_models.User{})
	return nil
}

func down0001(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	config := config.GetConfig()
	gormDB, err := gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	gormDB.Migrator().DropTable(&base_models.User{})
	return nil
}
