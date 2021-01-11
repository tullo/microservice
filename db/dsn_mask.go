package db

import "regexp"

var dsnMasker = regexp.MustCompile("(.)(?:.*)(.):(.)(?:.*)(.)@")

// maskDSN masks the DSN username and password for exposure in logs, etc.
//
// The regular expression takes care of masking any number of characters in
// the username and password part of the DSN, outputting only the first and
// last characters of each, with 4 asterisks in between.
func maskDSN(dsn string) string {
	return dsnMasker.ReplaceAllString(dsn, "$1****$2:$3****$4@")
}
