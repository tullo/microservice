package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Run takes migrations for a project and executes them against a database.
func Run(project string, db *sqlx.DB) error {
	fs, ok := migrations[project]
	if !ok {
		return errors.Errorf("Migrations for '%s' don't exist", project)
	}

	// run main migration
	if err := migrate(db, &fs, project, "migrations.sql"); err != nil {
		return err
	}

	// run service migrations
	for _, filename := range fs.Migrations() {
		if err := migrate(db, &fs, project, filename); err != nil {
			return err
		}
	}

	return nil
}

func migrate(db *sqlx.DB, fs *FS, project, filename string) error {
	log.Println("Running migrations from", filename)

	mig := migration{
		Project:        project,
		Filename:       filename,
		StatementIndex: -1,
		Status:         "",
	}

	// skip logs for main migrations table.
	useLog := (filename != "migrations.sql")
	if useLog {
		if err := db.Get(&mig, "select * from migrations where project=? and filename=?", mig.Project, mig.Filename); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("failed to select * from migrations where project=%v and filename=%v: %w", mig.Project, mig.Filename, err)
		}
		if mig.Status == "ok" {
			log.Println("Migrations already applied, skipping")

			return nil
		}
	}

	err := migrateUp(db, fs, &mig, useLog)
	if useLog {
		// log the migration status into the database
		set := func(fields []string) string {
			sql := make([]string, len(fields))
			for k, v := range fields {
				sql[k] = v + "=:" + v
			}

			return strings.Join(sql, ", ")
		}
		if _, err = db.NamedExec("replace into migrations set "+set(migrationFields), mig); err != nil {
			log.Println("Updating migration status failed:", err)
		}
	}

	return errors.Wrap(err, "migrateUp failed")
}

func migrateUp(db *sqlx.DB, fs *FS, mig *migration, useLog bool) error {
	stmts, err := statements(fs.ReadFile(mig.Filename))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error reading migration: %s", mig.Filename))
	}

	for idx, stmt := range stmts {
		// skip stmt if it has already been applied
		if idx >= mig.StatementIndex {
			mig.StatementIndex = idx
			if err := execQuery(db, idx, stmt, useLog); err != nil {
				mig.Status = err.Error()

				return err
			}
		}
	}
	mig.Status = "ok"

	return nil
}

func execQuery(db *sqlx.DB, idx int, query string, useLog bool) error {
	if useLog {
		log.Println()
		log.Println("-- Statement index:", idx)
		log.Println(query)
		log.Println()
	}

	if _, err := db.Exec(query); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to execute a query: %w", err)
	}

	return nil
}
