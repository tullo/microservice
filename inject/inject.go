package inject

import (
	"github.com/google/wire"
	"github.com/tullo/microservice/client"
	"github.com/tullo/microservice/db"
)

// Inject bundles wire.ProviderSet as a global variable for
// dependency injection use.
var Inject = wire.NewSet(
	db.Connect,
	Sonyflake,
	client.Inject,
)
