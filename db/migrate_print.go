package db

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func printQuery(idx int, query string) error {
	log.Println()
	log.Println("-- Statement index:", idx)
	log.Println(query)
	log.Println()

	return nil
}

func printMigrations(fs *FS, filename string) error {
	log.Println("Printing migrations from", filename)
	stmts, err := statements(fs.ReadFile(filename))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error reading migration: %s", filename))
	}

	for idx, stmt := range stmts {
		if err := printQuery(idx, stmt); err != nil {
			return err
		}
	}

	return nil
}

// Print prints main and service migrations.
func Print(project string) error {
	fs, ok := migrations[project]
	if !ok {
		return errors.Errorf("Migrations for '%s' don't exist", project)
	}

	// print main migration
	if err := printMigrations(&fs, "migrations.sql"); err != nil {
		return err
	}

	// print service migrations
	for _, filename := range fs.Migrations() {
		if err := printMigrations(&fs, filename); err != nil {
			return err
		}
	}

	return nil
}
