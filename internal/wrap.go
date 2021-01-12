package internal

import (
	"net/http"
	"strings"
)

// Get remote IP address.
func remoteIPAddr(r *http.Request) string {
	headers := [2]string{
		http.CanonicalHeaderKey("X-Forwarded-For"),
		http.CanonicalHeaderKey("X-Real-IP"),
	}
	for i := range headers {
		if addr := r.Header.Get(headers[i]); addr != "" {
			return strings.SplitN(addr, ", ", 2)[0]
		}
	}

	return strings.SplitN(r.RemoteAddr, ":", 2)[0]
}

// WrapWithIP wraps a http.Handler to inject the client IP into the context.
func WrapWithIP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := SetIPToContext(r.Context(), remoteIPAddr(r))
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
