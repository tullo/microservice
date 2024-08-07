//go:build wireinject
// +build wireinject

package stats

import (
	"context"

	"github.com/google/wire"

	"github.com/tullo/microservice/inject"
)

func New(ctx context.Context) (*Server, error) {
	wire.Build(
		NewFlusher,
		inject.Inject,
		wire.Struct(new(Server), "*"),
	)
	return nil, nil
}
