# 10 Assembling a binary

No JS side to compare against. Node runs your module directly, with no separate step where you
turn a library into a process. Go splits the two on purpose: everything before this module has
been a package with tests and no `func main`. This one is the wiring.

```go
func main() {
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
    defer stop()

    if err := run(ctx, os.Getenv, os.Stdout); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```

`main` is deliberately the only thing here with no test. It cannot be called from a test, because
there is nothing to call it *with*. So the real work moves into `run`, which takes `os.Getenv` and
`os.Stdout` as **parameters** instead of reaching for the globals directly. A test hands it a fake
`getenv` and a `bytes.Buffer` instead; production hands it the real thing. Same function, either
way.

## What run assembles

Every module before this one, wired together:

- **07 logging**: the `*slog.Logger` that writes JSON to `stdout`
- **09 service**: `service.New(store).Routes()`, the handler
- **08 context**: `ctx` is what makes the shutdown below possible at all

## Graceful shutdown

```go
<-ctx.Done()
shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
defer cancel()
srv.Shutdown(shutdownCtx)
```

`ctx.Done()` fires the moment `signal.NotifyContext` sees Ctrl-C. `Shutdown` then stops accepting
new connections and waits for the ones already in flight to finish, up to `shutdownTimeout`. A
request that was mid-response gets to complete instead of being cut off.

`Shutdown` takes its **own** context because `ctx` is already done by the time you call it.

## Run

```sh
go test ./cmd/
```

One test in here binds a real port instead of using `httptest`. This module is about the process
starting and answering, so the test proves both.
