package test

import (
	"os"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"

	"github.com/iegomez/smart-ac/internal/config"
	"github.com/iegomez/smart-ac/internal/migrations"
)

// GetConfig returns the test configuration.
func GetConfig() config.Config {
	log.SetLevel(log.ErrorLevel)

	var c config.Config

	c.PostgreSQL.DSN = "postgres://smart_ac:password@localhost/smart_ac_test?sslmode=disable"

	if v := os.Getenv("TEST_POSTGRES_DSN"); v != "" {
		c.PostgreSQL.DSN = v
	}

	return c
}

// MustResetDB re-applies all database migrations.
func MustResetDB(db *sqlx.DB) {
	m := &migrate.AssetMigrationSource{
		Asset:    migrations.Asset,
		AssetDir: migrations.AssetDir,
		Dir:      "",
	}
	if _, err := migrate.Exec(db.DB, "postgres", m, migrate.Down); err != nil {
		log.Fatal(err)
	}
	if _, err := migrate.Exec(db.DB, "postgres", m, migrate.Up); err != nil {
		log.Fatal(err)
	}
}
