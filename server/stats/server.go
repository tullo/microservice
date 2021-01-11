package stats

import (
	"context"

	"github.com/tullo/microservice/rpc/stats"
)

// Server implements stats.Stats.
type Server struct {
}

var _ stats.StatsService = &Server{}

func (svc *Server) Push(_ context.Context, _ *stats.PushRequest) (*stats.PushResponse, error) {
	panic("not implemented") // TODO: Implement
}
