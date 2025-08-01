package stats

import (
	"github.com/jmoiron/sqlx"
	"github.com/sony/sonyflake/2"
)

// Server implements stats.Stats.
type Server struct {
	db        *sqlx.DB
	sonyflake *sonyflake.Sonyflake
	flusher   *Flusher
}

// Shutdown is a cleanup hook after SIGTERM.
func (svc *Server) Shutdown() {
	<-svc.flusher.Done()
}
