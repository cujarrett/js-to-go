# 08 context.Context

```js
// JS: AbortController is the closest thing this has
const controller = new AbortController()
fetch(url, { signal: controller.signal })
setTimeout(() => controller.abort(), 50)
```

```go
// Go: cancellation is a parameter, not an option bag on one call
ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
defer cancel()
SlowOp(ctx, someDelay)
```

The difference that matters: in JS, the caller decides whether an operation is abortable by passing
`signal` or not. In Go, every function that can take time takes a `ctx context.Context` as its
**first** parameter, always - so cancellation is not something you opt a call into, it is something
you opt a call *out* of by ignoring `ctx.Done()`.

## select

```go
select {
case <-time.After(delay):
    return nil
case <-ctx.Done():
    return ctx.Err()
}
```

`select` waits on whichever channel is ready first. This is the one idiom behind almost every
timeout and cancellation path in Go: race the real work against `ctx.Done()`, and whichever
finishes first decides the return.

`ctx.Err()` is `context.Canceled` or `context.DeadlineExceeded` - a sentinel, exactly like
`ErrNotFound` in module 03. `errors.Is` is how the tests check which one fired.

## WithValue

No JS side here either - most JS code just passes an extra parameter for request-scoped data, and
the nearer analogue, Node's `AsyncLocalStorage`, is uncommon outside frameworks. Go's context
carries values for the same reason it carries cancellation: one parameter, threaded through every
layer, rather than a new argument added to every function along the way.

`context.WithValue` needs an unexported key type (`requestIDKey`), not a string - two packages
using the string `"requestID"` as a key would silently collide and overwrite each other's value.

## Run

```sh
go test ./context/
```
