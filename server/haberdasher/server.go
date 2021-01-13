package haberdasher

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/tullo/microservice/rpc/haberdasher"
)

// Server implements haberdasher.Haberdasher.
type Server struct {
	db *sqlx.DB
}

// Shutdown is a cleanup hook after SIGTERM.
func (*Server) Shutdown() {
}

// MakeHat produces a hat of mysterious, randomly-selected color!
func (*Server) MakeHat(context.Context, *haberdasher.Size) (*haberdasher.Hat, error) {
	return nil, nil
}

var _ haberdasher.HaberdasherService = &Server{}
