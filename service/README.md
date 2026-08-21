# 09 A real service

Everything so far - values, collections, errors, types, both JSON encoders, structured logging and context - in the shape every Go service in this workspace uses. Standard library only,
no framework - `net/http` is the framework.

```js
// Express
app.get("/servers/:name", (req, res) => res.json(find(req.params.name)))
```

```go
// Go 1.22+: method and path parameters live in the pattern
mux.HandleFunc("GET /servers/{name}", a.handleGet)
// inside the handler: r.PathValue("name")
```

Registering `GET /servers` also gets you a free 405 on `POST /servers`, which is why one of the
tests checks for it.

## Why Store is an interface

The handlers depend on a list of methods, not on a database. The test passes a `fakeStore` that
satisfies the same methods, so the whole API is tested with no network, no container, and no
mocking library. This is the Go habit: **accept interfaces, return structs.**

## httptest

`httptest.NewRequest` and `httptest.NewRecorder` run a handler directly, in-process. No port, no
server, no sleeping. Fast enough that you write more tests than you expected to.

## Run

```sh
go test ./service/
```

Once it is green, add `POST /servers` yourself: decode JSON from `r.Body`, return 201, and 400 on
bad input. Write the test first.
