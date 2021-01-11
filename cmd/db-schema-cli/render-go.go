package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"path"
	"strings"

	"github.com/pkg/errors"
)

var numericTypes map[string]string = map[string]string{
	"tinyint":   "int8",
	"smallint":  "int16",
	"mediumint": "int32", // this one would, technically, be int24 (3 bytes)
	"int":       "int32",
	"bigint":    "int64",
}

func isNumeric(column *Column) (string, bool) {
	val, ok := numericTypes[column.DataType]

	return val, ok
}

var simpleTypes map[string]string = map[string]string{
	"char":       "string",
	"varchar":    "string",
	"text":       "string",
	"longtext":   "string",
	"mediumtext": "string",
	"tinytext":   "string",
	"longblob":   "[]byte",
	"blob":       "[]byte",
	"varbinary":  "[]byte",
	"float":      "float32", // `float` and `double` are here ...
	"double":     "float64", // since they don't have unsigned modifiers
	"decimal":    "string",  // stored as string, \o/

}

func isSimple(column *Column) (string, bool) {
	val, ok := simpleTypes[column.DataType]

	return val, ok
}

type specialType struct {
	Import string
	Type   string
}

var specialTypes map[string]specialType = map[string]specialType{
	// `enum` and `set` aren't implemented
	// `year` isn't implemented
	"date":      specialType{"time", "*time.Time"},
	"datetime":  specialType{"time", "*time.Time"},
	"time":      specialType{"time", "*time.Time"},
	"timestamp": specialType{"time", "*time.Time"},
	"json":      specialType{"sqlxTypes github.com/jmoiron/sqlx/types", "sqlxTypes.JSONText"},
}

func isSpecial(column *Column) (specialType, bool) {
	val, ok := specialTypes[column.DataType]

	return val, ok
}

func contains(set []string, value string) bool {
	for _, v := range set {
		if v == value {
			return true
		}
	}
	return false
}

func resolveGoType(column *Column) (string, error) {
	if val, ok := isSimple(column); ok {
		return val, nil
	}
	if val, ok := isNumeric(column); ok {
		isUnsigned := strings.Contains(strings.ToLower(column.Type), "unsigned")
		if isUnsigned {
			return "u" + val, nil
		}
		return val, nil
	}
	if val, ok := isSpecial(column); ok {
		return val.Type, nil
	}
	return "", errors.Errorf("Unsupported SQL type: %s", column.DataType)
}

func renderGo(basePath string, service string, tables []*Table) error {

	imports := []string{}

	// Loop through tables/columns, return type error if any
	// This also builds the `imports` slice for codegen lower
	for _, table := range tables {
		for _, column := range table.Columns {
			if _, err := resolveGoType(column); err != nil {
				return err
			}
		}
	}
	buf := bytes.NewBuffer([]byte{})

	fmt.Fprintf(buf, "package %s\n", service)
	fmt.Fprintln(buf)

	// Print collected imports
	if len(imports) > 0 {
		fmt.Println("import (")
		for _, val := range imports {
			fmt.Printf("\t\"%s\"\n", val)
		}
		fmt.Println(")")
		fmt.Println()
	}

	for _, table := range tables {
		fields := []string{}
		primary := []string{}
		if table.Comment != "" {
			fmt.Println("//", table.Comment)
		}
		fmt.Printf("type %s struct {\n", camel(table.Name))
		for idx, column := range table.Columns {
			fields = append(fields, column.Name)
			if column.Key == "PRI" {
				primary = append(primary, column.Name)
			}

			if column.Comment != "" {
				if idx > 0 {
					fmt.Println()
				}
				fmt.Printf("	// %s\n", column.Comment)
			}
			columnType, _ := resolveGoType(column)
			fmt.Printf("    %s %s `db:\"%s\" json:\"-\"`\n", camel(column.Name), columnType, column.Name)
		}
		fmt.Println("}")
		fmt.Println()
		fmt.Printf("var %sFields = ", camel(table.Name))
		if len(fields) > 0 {
			fmt.Printf("[]string{\"%s\"}", strings.Join(fields, "\", \""))
		} else {
			fmt.Printf("[]string{}")
		}
		fmt.Println()
		fmt.Printf("var %sPrimaryFields = ", camel(table.Name))
		if len(primary) > 0 {
			fmt.Printf("[]string{\"%s\"}", strings.Join(primary, "\", \""))
		} else {
			fmt.Printf("[]string{}")
		}
		fmt.Println()
	}

	filename := path.Join(basePath, "types_gen.go")
	contents := buf.Bytes()

	formatted, err := format.Source(contents)
	if err != nil {
		// Fall back to unformatted source to inspect
		// the saved file for the error which occurred.
		formatted = contents
		log.Println("An error occurred while formatting the go source:", err)
		log.Println("Saving the unformatted code")
	}

	return ioutil.WriteFile(filename, formatted, 0600)
}
