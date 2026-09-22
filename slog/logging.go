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

// LogRequest records one handled request as structured attributes.
// "status" becomes a field a log pipeline can filter on, rather than text
// buried in a sentence.
func LogRequest(logger *slog.Logger, method, path string, status int) {
	// TODO
}

// WithRequestID returns a logger that stamps every later line with id, so the
// caller never passes it again. pino spells this logger.child({ requestId }).
func WithRequestID(logger *slog.Logger, id string) *slog.Logger {
	// TODO
	return nil
}
