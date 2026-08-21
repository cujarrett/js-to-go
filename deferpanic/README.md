# 13 defer, panic, recover

```js
// JS: finally always runs, same idea as defer
try {
  doWork()
} finally {
  cleanup()   // runs whether doWork returned or threw
}
```

```go
// Go: defer schedules a call for when the function returns, by any path
func WithCleanup(work func()) {
    defer cleanup()   // runs whether work() returns or panics
    work()
}
```

`defer` is `finally` turned into its own statement, attached to whatever call it sits in front of,
rather than wrapping a block. It runs in LIFO order if there is more than one, and it runs even if
the function panics on the way out - which is the whole reason finalizer and webhook code leans on
it: the resource gets released, the status gets written, no matter how the function ends.

## panic and recover

```js
// JS: throw / catch
try {
  risky()
} catch (e) {
  return { error: e.message }
}
```

```go
// Go: panic / recover, and recover only works inside a deferred func
func Safe(work func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    work()
    return nil
}
```

The one non-obvious rule: `recover()` only does anything when called **directly inside a deferred
function**. Call it anywhere else and it returns `nil` even during an active panic - it does not
reach up the stack to catch one the way `catch` does.

## Where you'll see this

A webhook or `Reconcile` that panics on one malformed object should not take the controller process
down with it. `Safe` here is that shape: catch, log, return an error - the reconcile gets requeued
instead of the whole manager crashing.

## Run

```sh
go test ./deferpanic/
```

`TestSafeTurnsPanicIntoError` prints a real panic trace until `Safe` recovers one - that crash is
the red test, not a mistake. Once `recover()` is in place, it turns into an ordinary `PASS`.
