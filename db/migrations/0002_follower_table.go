package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
	"Go-Rampup/config"
	"gorm.io/driver/postgres"
	base_models "Go-Rampup/db/models/base"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(up0002, down0002)
}

func up0002(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	config := config.GetConfig()
	gormDB, err := gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	gormDB.Migrator().CreateTable(&base_models.Follower{})
	return nil
}

func down0002(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	config := config.GetConfig()
	gormDB, err := gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	gormDB.Migrator().DropTable(&base_models.Follower{})
	return nil
}
