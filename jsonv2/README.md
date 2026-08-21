# 06 encoding/json/v2

No JS side this time - this is a Go-only story, v1 of its own standard library against v2.

`encoding/json/v2` is real, accepted, and still experimental: it does not build without
`GOEXPERIMENT=jsonv2` set, which is why `just test 06` and `just ci` set it for you. Running plain
`go test ./jsonv2/` without that variable fails to compile, on purpose - it is Go's usual signal
that something is not yet a promise.

## The bug v2 fixes

```go
type Record struct {
    LastSynced time.Time `json:"lastSynced,omitempty"`
}
```

In v1, `omitempty` never omits a **struct** field, zero value or not - it only recognises nil
pointers, empty maps/slices, and zero scalars. A zero `time.Time` serialises as
`"0001-01-01T00:00:00Z"` regardless of what you write here. This has surprised nearly everyone who
has shipped a Go API with an optional timestamp.

```go
LastSynced time.Time `json:"lastSynced,omitzero"`
```

`omitzero` asks the type itself: if it has an `IsZero() bool` method, that decides; otherwise it
compares against the type's zero value directly. `time.Time` already implements `IsZero`, so this
one tag change fixes the bug with no other code involved.

## Run

```sh
just test 06
```

Plain `go test ./jsonv2/` intentionally fails to compile - that is the GOEXPERIMENT gate, not a
mistake in your code.
