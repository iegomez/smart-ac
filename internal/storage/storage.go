package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"

	"github.com/iegomez/smart-ac/internal/config"
	"github.com/iegomez/smart-ac/internal/migrations"
)

var (
	// The secret used to generate JWT tokens.
	jwtsecret []byte
	// db holds the PostgreSQL connection pool.
	db *sqlx.DB
)

// DB returns the database object.
func DB() *sqlx.DB {
	return db
}

// Transaction wraps the given function in a transaction. In case the given
// functions returns an error, the transaction will be rolled back.
func Transaction(f func(tx sqlx.Ext) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return errors.Wrap(err, "storage: begin transaction error")
	}

	err = f(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Wrap(rbErr, "storage: transaction rollback error")
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "storage: transaction commit error")
	}
	return nil
}

// Setup configures the storage package.
func Setup(c config.Config) error {
	log.Info("storage: setting up storage package")

	jwtsecret = []byte(c.ExternalAPI.JWTSecret)

	log.Info("storage: connecting to PostgreSQL database")
	d, err := sqlx.Open("postgres", c.PostgreSQL.DSN)
	if err != nil {
		return errors.Wrap(err, "storage: PostgreSQL connection error")
	}
	for {
		if err := d.Ping(); err != nil {
			log.WithError(err).Warning("storage: ping PostgreSQL database error, will retry in 2s")
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	db = d

	if c.PostgreSQL.Automigrate {
		log.Info("storage: applying PostgreSQL data migrations")
		m := &migrate.AssetMigrationSource{
			Asset:    migrations.Asset,
			AssetDir: migrations.AssetDir,
			Dir:      "",
		}
		n, err := migrate.Exec(db.DB, "postgres", m, migrate.Up)
		if err != nil {
			return errors.Wrap(err, "storage: applying PostgreSQL data migrations error")
		}
		log.WithField("count", n).Info("storage: PostgreSQL data migrations applied")
	}

	return nil
}
