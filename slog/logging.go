// Package logging is structured logging the way every Go service in this
// workspace does it: log/slog, one JSON object per line, to stdout.
package logging

import (
	"io"
	"log/slog"
)

// New returns a logger that writes JSON lines to w.
func New(w io.Writer) *slog.Logger {
	// TODO
	return nil
}

// LogRequest records one handled request as structured attributes, not a
// formatted string - so "status" is a field a log pipeline can filter on,
// not a substring inside a sentence.
func LogRequest(logger *slog.Logger, method, path string, status int) {
	// TODO
}

// WithRequestID returns a logger that stamps every future line with id,
// without the caller passing it again. Node's pino calls the same idea
// logger.child({ requestId }); slog calls it With.
func WithRequestID(logger *slog.Logger, id string) *slog.Logger {
	// TODO
	return nil
}
