package db

import "strings"

// cleanDSN will append default options to the DSN (Data Source Name),
// if they are not already provided.
// If you specify ?loc=UTC to your microservice DSN environment,
// that option will be used, instead of the default loc=Local.
func cleanDSN(dsn string) string {
	// collation=utf8mb4_general_ci
	// - set the default collation
	// (utf8mb4 is a given these days if you want emojis)
	dsn = addOptionToDSN(dsn, "?", "?")
	dsn = addOptionToDSN(dsn, "collation=", "&collation=utf8_general_ci")
	dsn = addOptionToDSN(dsn, "parseTime=", "&parseTime=true") // for decoding date/datetime values into a time.Time
	dsn = addOptionToDSN(dsn, "loc=", "&loc=Local")            // location for time.Time values
	dsn = strings.Replace(dsn, "?&", "?", 1)

	return dsn
}

func addOptionToDSN(dsn, match, option string) string {
	if !strings.Contains(dsn, match) {
		dsn += option
	}

	return dsn
}
