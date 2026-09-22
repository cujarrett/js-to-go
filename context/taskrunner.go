// Package taskrunner is context.Context: cancellation, deadlines, and
// request-scoped values, the thing that flows through every layer of a real
// Go service the way an AbortSignal flows through a fetch call in JS.
package taskrunner

import (
	"context"
	"time"
)

// SlowOp simulates work that takes delay to finish. It must stop early and
// return ctx.Err() the moment ctx is done, rather than waiting delay out.
func SlowOp(ctx context.Context, delay time.Duration) error {
	// TODO: select on time.After(delay) and ctx.Done()
	return nil
}

type requestIDKey struct{}

// WithRequestID returns a copy of ctx carrying id. JS has no true equivalent -
// most code just passes an extra parameter, and Node's AsyncLocalStorage,
// the closer analogue, is rare in ordinary handlers. Go leans on ctx because
// cancellation and request-scoped values travel through the same parameter
// by convention, so nothing extra has to be threaded through every call.
func WithRequestID(ctx context.Context, id string) context.Context {
	// TODO
	return ctx
}

// RequestID reads the id WithRequestID stored, if any.
func RequestID(ctx context.Context) (string, bool) {
	// TODO
	return "", false
}
