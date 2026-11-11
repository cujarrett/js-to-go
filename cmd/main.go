// Command server is every module before it, wired into one runnable process the
// way this workspace wires every Go service: env-configured, structured logging,
// graceful shutdown. func main is deliberately the only thing not under test -
// run does the real work, so it is the only thing that has to be.
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"js-to-go/service"
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
	logger := slog.New(slog.NewJSONHandler(stdout, nil))

	port := getenv("PORT")
	if port == "" {
		port = "8080"
	}

	store := &memStore{servers: []service.Server{
		{Name: "web-1", Slot: "demo1", Active: true},
	}}
	srv := &http.Server{Addr: ":" + port, Handler: service.New(store).Routes()}

	serveErr := make(chan error, 1)
	go func() {
		err := srv.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		serveErr <- err
	}()
	logger.Info("listening", "port", port)

	select {
	case err := <-serveErr:
		return err
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}
	return <-serveErr
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
