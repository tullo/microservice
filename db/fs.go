package db

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// FS represents a mapping between filename => contents.
type FS map[string]string

// Migrations returns list of SQL files to execute.
func (fs FS) Migrations() []string {
	var result []string
	for filename, contents := range fs {
		// skip empty files
		if contents == "" {
			continue
		}
		if matched, _ := filepath.Match("*.up.sql", filename); matched {
			result = append(result, filename)
		}
	}
	sort.Strings(result)

	return result
}

// ReadFile returns decoded file contents from FS.
func (fs FS) ReadFile(filename string) ([]byte, error) {
	if val, ok := fs[filename]; ok {
		bs, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return nil, fmt.Errorf("failed to decode: %w", err)
		}

		return bs, nil
	}

	return nil, os.ErrNotExist
}
