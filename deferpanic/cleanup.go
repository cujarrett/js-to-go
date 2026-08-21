// Package cleanup is defer, panic and recover - the pattern behind finalizer
// and webhook handler code, where cleanup must run whether a function returns
// normally or blows up partway through.
package cleanup


// Log records what happened, in order - the test reads this to prove defer
// ran last, not whether you can see console output.
type Log struct {
	lines []string
}

func (l *Log) add(line string) { l.lines = append(l.lines, line) }

// WithCleanup runs work, and afterward always logs "cleanup" - whether work
// returns normally or panics. This is the finalizer shape: release a resource
// no matter how the function that used it ends.
func WithCleanup(log *Log, work func()) {
	// TODO: defer a func that logs "cleanup", then call work()
}

// Safe runs work and turns a panic into a returned error instead of letting it
// crash the process - the shape of a webhook or reconcile handler that must
// not take the whole controller down over one bad object.
func Safe(work func()) (err error) {
	// TODO: defer a func that recovers, and on a non-nil recovered value sets
	// err to fmt.Errorf("recovered: %v", ...)
	work()
	return nil
}
