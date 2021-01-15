package stats

import (
	"context"
	"fmt"
	"time"

	"github.com/tullo/microservice/internal"
	"github.com/tullo/microservice/rpc/stats"
)

// Keep returning a single object to avoid allocations.
var pushResponseDefault = new(stats.PushResponse)

// Push a record to the incoming log table.
func (svc *Server) Push(ctx context.Context, r *stats.PushRequest) (*stats.PushResponse, error) {
	ctx = internal.ContextWithoutCancel(ctx)

	var err error
	if err := validate(r); err != nil {
		return nil, err
	}

	var row Incoming
	row.SetStamp(time.Now())
	row.ID, err = svc.sonyflake.NextID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate sonyflake id: %w", err)
	}
	row.Property = r.Property
	row.PropertySection = r.Section
	row.PropertyID = r.Id
	row.RemoteIP = internal.GetIPFromContext(ctx)

	return pushResponseDefault, svc.flusher.Push(&row)
}

func validate(r *stats.PushRequest) error {
	if r.Property == "" {
		return errMissingProperty
	}
	if r.Property != "news" {
		return errInvalidProperty
	}
	if r.Id < 1 {
		return errMissingID
	}
	if r.Section < 1 {
		return errMissingSection
	}

	return nil
}
