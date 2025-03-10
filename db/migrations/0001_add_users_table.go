package migrations

import (
	"Go-Rampup/config"
	"context"
	"database/sql"
	"gorm.io/gorm"
	models "Go-Rampup/db/models"
	"gorm.io/driver/postgres"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddUsersTable, downAddUsersTable)
}

func upAddUsersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	config := config.GetConfig()
	gormDB, err := gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	gormDB.Migrator().CreateTable(&models.User{})
	return nil
}

func downAddUsersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	config := config.GetConfig()
	gormDB, err := gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	gormDB.Migrator().DropTable(&models.User{})
	return nil
}
