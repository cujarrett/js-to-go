# learning-go

A JS-to-Go track, then a second track aimed straight at reading and writing Kubernetes
controllers. Each module is a package of stubbed functions and a test file that fails until you
write them.

The loop is always the same: **run the tests, read the failure, write the code, run again.** No
tutorial to follow, no output to eyeball. The tests tell you when you are done.

Written for someone comfortable in JavaScript, so each README compares the two where it actually
helps and stays quiet where it does not. Zero external dependencies, on purpose - every module
imports only the standard library.

Order lives here, not in the folder names - a table entry moves for the cost of one line, a
renumbered directory used to cost a rename across every cross-reference in the repo.

## Order

### Track A - Go, from a JS lens

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

`service` uses everything before it except `cmd`, which wires `service` itself into a runnable
process.

### Track B - reading real controllers

Everything a `kubebuilder`-scaffolded controller uses that Track A does not touch yet.

| Module | Idea | JS you already know |
|---|---|---|
| [embedding](./embedding) | struct embedding, method promotion | `extends`, and why Go has no such thing |
| [generics](./generics) | type parameters, constraints | nothing - JS never asks |
| [deferpanic](./deferpanic) | cleanup that runs no matter how a function ends | `try / finally`, `throw / catch` |

Do every module in the order each table lists it, top to bottom - within a track, and Track A
before Track B.

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

## After Track B

- **Concurrency** - goroutines, channels, worker pools. Deliberately not here: it is what Go is
  famous for and the least of what you need day to day. `context` covers cancellation on its own,
  since it shows up in ordinary code long before a goroutine does.
- **Kubernetes controllers, for real** - [The Kubebuilder Book](https://book.kubebuilder.io/)
  CronJob tutorial, then a controller of your own that uses `Owns()`.
- **Idiomatic Go** - [Effective Go](https://go.dev/doc/effective_go) and
  [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments), best read once the syntax is
  automatic.
