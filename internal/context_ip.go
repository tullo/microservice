package internal

import (
	"context"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// ipAddressCtxKey is used to store/retrieve a string value from a context.Context.
const ipAddressCtxKey ctxKey = 1

// SetIPToContext sets IP value to ctx.
func SetIPToContext(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, ipAddressCtxKey, ip)
}

// GetIPFromContext gets IP value from ctx.
func GetIPFromContext(ctx context.Context) string {
	if ip, ok := ctx.Value(ipAddressCtxKey).(string); ok {
		return ip
	}

	return ""
}
