package haberdasher

import (
	"github.com/jmoiron/sqlx"
	"github.com/sony/sonyflake"
)

// Server implements haberdasher.Haberdasher.
type Server struct {
	db        *sqlx.DB
	sonyflake *sonyflake.Sonyflake
}

// Shutdown is a cleanup hook after SIGTERM.
func (*Server) Shutdown() {
}
