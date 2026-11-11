# 06 encoding/json/v2

No JS side this time. This is Go against itself, v1 of its own standard library against v2.

Import it as `"encoding/json/v2"`. The package is still called `json`, and `json.Marshal(v)` works
the same as in v1.

- Go 1.26: v2 builds only with `GOEXPERIMENT=jsonv2` set. `just test jsonv2` and `just ci` set it.
- Go 1.27 and later: plain `go test ./jsonv2/` builds too.

## The bug v2 fixes

```go
type Record struct {
    LastSynced time.Time `json:"lastSynced,omitempty"`
}
```

In v1, `omitempty` never omits a **struct** field, zero value or not. It only recognises nil
pointers, empty maps/slices, and zero scalars. A zero `time.Time` serialises as
`"0001-01-01T00:00:00Z"` regardless of what you write here. This has surprised nearly everyone who
has shipped a Go API with an optional timestamp.

```go
LastSynced time.Time `json:"lastSynced,omitzero"`
```

`omitzero` asks the type itself: if it has an `IsZero() bool` method, that decides; otherwise it
compares against the type's zero value directly. `time.Time` already implements `IsZero`, so this
one tag change fixes the bug with no other code involved.

`omitzero` came with the v2 work. Go 1.24 added it to v1 as well.

## Run

```sh
just test jsonv2
```
