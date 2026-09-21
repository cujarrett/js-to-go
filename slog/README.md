# 07 Structured logging

`console.log` writes a sentence. A structured logger writes a record of fields a machine can filter
and group on.

```js
// JS: pino, one of the closer analogues
logger.info({ method: "GET", path: "/servers", status: 200 }, "request")
```

```go
// Go: log/slog, alternating key then value instead of one object
logger.Info("request", "method", "GET", "path", "/servers", "status", 200)
```

Both end up as one JSON object. slog takes `any` pairs instead of a map so logging a request in the
hot path costs no allocation when nothing is listening. The fields are only built into a record if
a handler is actually attached.

## New

`slog.New(slog.NewJSONHandler(w, nil))` is the line every service here starts with: JSON lines, to
whatever `io.Writer` fits, `os.Stdout` in production and a `bytes.Buffer` in a test.

## With

```js
const scoped = logger.child({ requestId })
```

```go
scoped := logger.With("request_id", id)
```

Same idea under a different name: a logger that remembers some fields so every call site after it
does not have to repeat them. `scoped` is a new value. The logger passed into `With` is unchanged,
which is what the third test checks.

## Run

```sh
go test ./slog/
```
