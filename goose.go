package main

import (
	"Go-Rampup/config"
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	_ "Go-Rampup/db/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	config := config.GetConfig()
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	dbDSN := config.DB.DSN
	driver := config.DB.Driver
	migrationsDir := config.DB.MigrationsDir

	goose.SetDialect(driver)
	db, err := sql.Open(driver, dbDSN)
	if err != nil {
		log.Fatalf("-dsn=%q: %v\n", dbDSN, err)
	}
	if err := goose.RunContext(context.Background(), args[0], db, migrationsDir, args[1:]...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}
