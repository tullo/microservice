package client

import (
	"github.com/google/wire"
	"github.com/tullo/microservice/client/haberdasher"
	"github.com/tullo/microservice/client/stats"
)

// Inject produces a wire.ProviderSet with our RPC clients.
var Inject = wire.NewSet(
	haberdasher.New,
	stats.New,
)
