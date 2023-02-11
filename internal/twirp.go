package internal

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/twitchtv/twirp"
	"go.elastic.co/apm/v2"
)

// NewServerHooks provides an error logging hook with Elastic APM.
func NewServerHooks() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		Error: func(ctx context.Context, err twirp.Error) context.Context {
			apm.CaptureError(ctx, err).Send()

			return ctx
		},
	}
}

// NewLoggingHooks creates a new twirp.ServerHooks which logs requests as they
// are routed to Twirp, and logs responses (including response time) when they
// are sent.
//
// Examples: NewLoggingHooks(os.Stderr) or (local/remote) log writer.
func NewLoggingHooks(w io.Writer) *twirp.ServerHooks {
	return &twirp.ServerHooks{
		Error: func(ctx context.Context, err twirp.Error) context.Context {
			apm.CaptureError(ctx, err).Send()

			return ctx
		},
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			startTime := time.Now()
			ctx = context.WithValue(ctx, hooksCtxKey, startTime)

			return ctx, nil
		},
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			svc, _ := twirp.ServiceName(ctx)
			meth, _ := twirp.MethodName(ctx)
			fmt.Fprintf(w, "received req svc=%q method=%q\n", svc, meth)

			return ctx, nil
		},
		ResponseSent: func(ctx context.Context) {
			startTime := ctx.Value(hooksCtxKey).(time.Time)
			svc, _ := twirp.ServiceName(ctx)
			meth, _ := twirp.MethodName(ctx)
			fmt.Fprintf(w, "response sent svc=%q method=%q time=%q\n", svc, meth, time.Since(startTime))
		},
	}
}
