# learning-go

Six modules that take you from "I can read Go" to "Go is not the hard part". Each is a package of
stubbed functions and a test file that fails until you write them.

The loop is always the same: **run the tests, read the failure, write the code, run again.** No
tutorial to follow, no output to eyeball. The tests tell you when you are done.

Written for someone comfortable in JavaScript, so each README compares the two where it actually
helps and stays quiet where it does not.

## Order

| Module | Idea | JS you already know |
|---|---|---|
| [01 values](./01_values) | pointers, copies, nil | objects are always by reference |
| [02 collections](./02_collections) | slices, maps, sets | arrays and plain objects |
| [03 errors](./03_errors) | errors as values | try / catch |
| [04 types](./04_types) | methods, receivers, interfaces | classes and duck typing |
| [05 json](./05_json) | struct tags, encoding | JSON.stringify |
| [06 service](./06_service) | an HTTP API, stdlib only | Express, minus Express |

Do them in order, one per sitting. Module 06 uses everything before it.

## Running

```sh
go test ./01_values/          # one module
go test ./...                 # everything
go test -race ./...           # how CI runs it
```

A red test naming the thing you have not written yet is the point. Nothing here should be read
without being run.

## How to get the most from it

- **Type the code, do not paste it.** Most of the value is muscle memory.
- **Predict the failure before running.** Being wrong is the signal worth having.
- **Break each module deliberately** once it is green: remove an `&`, change `%w` to `%v`, drop a
  struct tag. The compiler and the tests are better tutors than any explanation.
- **`go doc` beats a search engine.** `go doc net/http.ServeMux`, offline and exact.

## After module 06

- **Concurrency** - goroutines, channels, `context`. Deliberately not here: it is what Go is
  famous for and the least of what you need day to day.
- **Kubernetes controllers** - [The Kubebuilder Book](https://book.kubebuilder.io/) CronJob
  tutorial, then a controller that uses `Owns()`.
- **Idiomatic Go** - [Effective Go](https://go.dev/doc/effective_go) and
  [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments), best read once the syntax is
  automatic.
