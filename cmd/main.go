// Command server is every module before it, wired into one runnable process the
// way this workspace wires every Go service: env-configured, structured logging,
// graceful shutdown. func main is deliberately the only thing not under test -
// run does the real work, so it is the only thing that has to be.
package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"

	"learning-go/service"
)

// memStore is the one dependency service.API needs. A real binary would back
// this with a database; this exercise is about assembly, not storage.
type memStore struct {
	servers []service.Server
}

func (m *memStore) Get(name string) (service.Server, bool) {
	for _, s := range m.servers {
		if s.Name == name {
			return s, true
		}
	}
	return service.Server{}, false
}

func (m *memStore) List() []service.Server { return m.servers }

// run builds the logger and the server, starts listening, and blocks until ctx
// is done - then shuts down cleanly instead of dropping connections. getenv and
// stdout are parameters rather than os.Getenv/os.Stdout directly, which is what
// makes this function testable without a real environment or a real terminal.
func run(ctx context.Context, getenv func(string) string, stdout io.Writer) error {
	// TODO: build a *slog.Logger writing JSON to stdout - see module 07

	// TODO: read PORT from getenv, default "8080" if unset

	// TODO: build a memStore, then service.New(store).Routes() for the handler

	// TODO: construct an *http.Server on ":"+port with the routes above

	// TODO: start it in a goroutine, log the port it is listening on

	// TODO: block on <-ctx.Done(), then Shutdown with a fresh context bounded by
	// a few seconds - one of the exact things module 08 exists for

	return nil
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := run(ctx, os.Getenv, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// shutdownTimeout bounds how long a graceful shutdown waits for connections to
// drain before giving up.
const shutdownTimeout = 5 * time.Second
