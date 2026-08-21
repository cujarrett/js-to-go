# 12 Generics

```js
// JS never asks: any function already takes any type
const names = (items) => items.map(i => i.name)
```

```go
// Go: say what the function needs from T, once, and the compiler checks every call
func Names[T Named](items []T) []string { ... }

type Named interface{ Name() string }
```

JS's freedom is the whole story on that side - a function works on whatever you hand it, and if it
does not have `.name`, you find out at runtime. Go asks you to say what a type parameter needs
**once**, as a constraint, and then checks every call site against it before the program runs at
all.

`Names[Server]` and `Names[Pod]` are two different types, and `Names` itself does not change
between them - same idea as module 04's `CopyAll(dst Store, ...)`, one level more general. An
interface constrains a **value** you already have; a constraint like `Named` here constrains a
**type parameter**, before any value exists yet.

## Where you'll actually see this

`client-go`'s newer informers and listers, and `controller-runtime`'s typed client
(`client.ObjectList`, `runtime.Object`-adjacent generic helpers), are written this way so one
implementation serves every Kubernetes type without a copy per type or a cast back from `any`.

## Set[T]

```go
type Set[T comparable] struct {
    items map[T]bool
}
```

Same trick as module 02's `map[string]bool`, generalised: `comparable` is a built-in constraint
meaning "can be a map key" - which is exactly the requirement a set has, whatever `T` turns out to
be.

## Run

```sh
go test ./generics/
```
