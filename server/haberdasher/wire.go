//+build wireinject

package haberdasher

import (
	"context"

	"github.com/google/wire"

	"github.com/tullo/microservice/inject"
)

func New(ctx context.Context) (*Server, error) {
	wire.Build(
		inject.Inject,
		wire.Struct(new(Server), "*"),
	)
	return nil, nil
}
