package db

import "github.com/jmoiron/sqlx"

// Run takes migrations for a project and executes them against a database.
func Run(project string, db *sqlx.DB) error {
	return nil
}
