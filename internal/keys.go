package internal

// contextKey represents the type of value for the context keys.
type contextKey int

const (
	ipAddressCtxKey contextKey = 1 + iota
	hooksCtxKey
)
