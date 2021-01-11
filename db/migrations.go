package db

import (
	"regexp"
	"strings"
)

type migration struct {
	Project        string `db:"project"`
	Filename       string `db:"filename"`
	StatementIndex int    `db:"statement_index"`
	Status         string `db:"status"`
}

var migrationFields = []string{
	"project",         // denotes the migration group (e.g. stats)
	"filename",        // migration from that group (2021-01-11-092610-description-here.up.sql)
	"statement_index", // sequential statement index from a migration
	"status",          // "ok", or the error produced from a failing migration
}

// Read migrations and split them into statements.
func statements(contents []byte, err error) ([]string, error) {
	result := []string{}
	if err != nil {
		return result, err
	}

	stmts := regexp.MustCompilePOSIX(";$").Split(string(contents), -1)
	for _, stmt := range stmts {
		stmt = strings.TrimSpace(stmt)
		if stmt != "" {
			result = append(result, stmt)
		}
	}

	return result, nil
}
