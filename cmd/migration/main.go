package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/greeflas/racing-engine-backend/migrations"

	"github.com/uptrace/bun/migrate"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	connStr := "postgres://postgres:pass@localhost:5433/racing_engine?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connStr)))

	db := bun.NewDB(sqldb, pgdialect.New())
	defer db.Close()

	migrator := migrate.NewMigrator(db, migrations.Migrations)

	ctx := context.Background()
	if err := migrator.Init(ctx); err != nil {
		log.Printf("Init failed: %v", err)
		return
	}

	if len(os.Args) < 2 {
		log.Printf("No comand name provided. Usage: %s create|migrate", os.Args[0])
		return
	}

	switch os.Args[1] {
	case "create":
		if len(os.Args) < 3 {
			log.Printf("No migration name provided. Usage: %s %s migration_name", os.Args[0], os.Args[1])
			return
		}

		migrationName := os.Args[2]

		mf, err := migrator.CreateGoMigration(ctx, migrationName)
		if err != nil {
			log.Printf("Creating failed: %v", err)
			return
		}

		log.Printf("Migration %s created successfully.\n", mf.Name)
	case "migrate":
		group, err := migrator.Migrate(ctx)
		if err != nil {
			log.Printf("Migration failed: %v", err)
			return
		}

		if group.IsZero() {
			log.Println("No new migrations found.")
		} else {
			log.Printf("Migrated successfully to %s.\n", group)
		}
	}
}
