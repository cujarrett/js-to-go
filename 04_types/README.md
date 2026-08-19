# 04 Types

A method is a function with a receiver. The receiver is `this`, and you choose whether it is a
copy or the real thing.

```js
class Counter {
  inc() {
    this.n++   // always the real thing
  }
}
```

```go
func (c *Counter) Inc() {
    c.n++   // the real thing
}
```

`(c *Counter)` before the name is the receiver - the same choice as module 01, applied to
`this`. And as in module 01, Go's other option has no JS equivalent:

```go
func (c Counter) IncByValue() {
    c.n++   // a copy, so this is a no-op
}
```

Rule of thumb: **if a method changes anything, use a pointer receiver.** Controllers use pointer
receivers everywhere.

## Interfaces

A Go interface is a list of methods. Nothing declares that it implements one - if the methods
exist, it fits. This is duck typing, checked at compile time.

```js
// JS: duck typing, discovered at runtime
const copyAll = (dst, pairs) => {
  for (const [k, v] of pairs) {
    dst.put(k, v)
  }
}
```

```go
// Go: duck typing, proven at compile time
func CopyAll(dst Store, pairs map[string]string) {
    for k, v := range pairs {
        dst.Put(k, v)
    }
}
```

Only the signature differs. `dst Store` names an interface, so the compiler checks at build time
what JS finds out when `dst.put` is missing at runtime.

`var s Store = NewMemStore()` in the test is the compile-time proof. Get a method signature wrong
and the test file will not build - which is the error you want.

Small interfaces are the Go habit: one or two methods, defined where they are used rather than
next to the type that satisfies them. `client.Object` in controller-runtime is exactly this.

## Run

```sh
go test ./04_types/
```
