# js-to-go

Ten modules that take you from "I can read Go" to "Go is not the hard part." Each is a package of
stubbed functions and a test file that fails until you write them.

The loop is always the same: **run the tests, read the failure, write the code, run again.** No
tutorial to follow, no output to eyeball. The tests tell you when you are done.

Written for someone comfortable in JavaScript, so each README compares the two where it actually
helps and stays quiet where it does not. Zero external dependencies, on purpose - every module
imports only the standard library.

## Order

| Module | Idea | JS you already know |
|---|---|---|
| [values](./values) | pointers, copies, nil | objects are always by reference |
| [collections](./collections) | slices, maps, sets | arrays and plain objects |
| [errors](./errors) | errors as values | try / catch |
| [types](./types) | methods, receivers, interfaces | classes and duck typing |
| [json](./json) | struct tags, encoding | JSON.stringify |
| [jsonv2](./jsonv2) | a real stdlib package, still experimental | no JS side - see the README |
| [slog](./slog) | structured logging | pino / bunyan |
| [context](./context) | cancellation, deadlines, request values | AbortController |
| [service](./service) | an HTTP API, stdlib only | Express, minus Express |
| [cmd](./cmd) | wiring a binary, graceful shutdown | no JS side - see the README |

Do them in order, one per sitting. `service` uses everything before it except `cmd`, which wires
`service` itself into a runnable process.

## Running

```sh
go test ./values/          # one module
just ci                    # everything, the way CI runs it
```

`just ci` is what to use for "everything" - `jsonv2` needs `GOEXPERIMENT=jsonv2` to even compile,
which plain `go test ./...` does not know to set, and `just` does.

A red test naming the thing you have not written yet is the point. Nothing here should be read
without being run.

## How to get the most from it

- **Type the code, do not paste it.** Most of the value is muscle memory.
- **Predict the failure before running.** Being wrong is the signal worth having.
- **Break each module deliberately** once it is green: remove an `&`, change `%w` to `%v`, drop a
  struct tag. The compiler and the tests are better tutors than any explanation.
- **`go doc` beats a search engine.** `go doc net/http.ServeMux`, offline and exact.

## After this repo

- **[go-to-controllers](https://github.com/cujarrett/go-to-controllers)** - the Go idioms a real
  Kubernetes controller leans on that a JS comparison never needed to touch: struct embedding,
  generics, defer/panic/recover.
- **Concurrency** - goroutines, channels, worker pools. Deliberately not here: it is what Go is
  famous for and the least of what you need day to day. `context` covers cancellation on its own,
  since it shows up in ordinary code long before a goroutine does.
- **Idiomatic Go** - [Effective Go](https://go.dev/doc/effective_go) and
  [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments), best read once the syntax is
  automatic.
