package stats

import "errors"

var (
	errFlusherDisabled = errors.New("Flusher is disabled, shutting down")
	errInvalidProperty = errors.New("invalid property")
	errMissingID       = errors.New("missing id")
	errMissingProperty = errors.New("missing property")
	errMissingSection  = errors.New("missing section")
)
