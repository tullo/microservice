package haberdasher

// Hat generated for db table `hat`.
//
// Hat is a piece of headwear made by a Haberdasher.
type Hat struct {
	// Tracking ID
	ID uint64 `db:"id" json:"-"`

	// The size of a hat in centimeters.
	Size uint32 `db:"size" json:"-"`

	// The name of a hat.
	Color string `db:"color" json:"-"`

	// The color of a hat.
	Name string `db:"name" json:"-"`
}

// HatTable is the name of the table in the DB.
const HatTable = "`hat`"

// HatFields are all the field names in the DB table.
var HatFields = []string{"id", "size", "color", "name"}

// HatPrimaryFields are the primary key fields in the DB table.
var HatPrimaryFields = []string{"id"}

// Migrations generated for db table `migrations`.
type Migrations struct {
	// Microservice or project name.
	Project string `db:"project" json:"-"`

	// yyyy-mm-dd-HHMMSS-filename.up.sql.
	Filename string `db:"filename" json:"-"`

	// Statement number from SQL file.
	StatementIndex int32 `db:"statement_index" json:"-"`

	// ok or full error message.
	Status string `db:"status" json:"-"`
}

// MigrationsTable is the name of the table in the DB.
const MigrationsTable = "`migrations`"

// MigrationsFields are all the field names in the DB table.
var MigrationsFields = []string{"project", "filename", "statement_index", "status"}

// MigrationsPrimaryFields are the primary key fields in the DB table.
var MigrationsPrimaryFields = []string{"project", "filename"}
