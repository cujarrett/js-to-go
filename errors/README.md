# 03 Errors

Go has no `throw`. A failure is a value you return, and the caller decides what to do with it.

```js
// JS
try {
  const v = get(key)
  return v
} catch (e) {
  if (e.code === "ENOENT") return ""
  throw e
}
```

```go
// Go
v, err := get(key)
if errors.Is(err, ErrNotFound) {
    return "", nil
}
if err != nil {
    return "", err
}
```

Same logic, spelled out as ordinary code. This is why controllers read the way they do:

```go
if apierrors.IsNotFound(err) { continue }   // the specific case first
if err != nil { return err }                // then the general one
```

## Three tools

- **Sentinel** - `var ErrNotFound = errors.New("not found")`. Compare with `errors.Is`.
- **Wrapping** - `fmt.Errorf("describing %s: %w", key, err)`. `%w` keeps the original reachable.
  `%v` flattens it to text and breaks `errors.Is`.
- **Custom type** - a struct with fields, pulled back out with `errors.As` when the caller needs
  the detail and not just the fact. `errors.AsType[*MissingKeyError](err)` (Go 1.26+) does the same
  walk in one call instead of three lines - the type parameter replaces the `var` and the `&`.

## Run

```sh
go test ./errors/
```

Then change one `%w` to `%v` and watch a test fail. That is the whole lesson.
