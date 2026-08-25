# 02 Collections

Slices are arrays. Maps are plain objects. The differences are worth knowing.

| JavaScript | Go |
|---|---|
| `const xs = []` | `var xs []string` (nil, but you can append) |
| `xs.push(x)` | `xs = append(xs, x)` - it returns a new slice, you must reassign |
| `xs.map(f)` | a `for` loop. Go has no map/filter/reduce, and does not miss them |
| `const m = {}` | `m := map[string]int{}` |
| `m[k] ?? 0` | `m[k]` - a missing key already gives the zero value |
| `k in m` | `v, ok := m[k]` |
| `new Set()` | `map[string]bool` |

## Two that bite

```go
for _, s := range servers {
    s.Active = true   // does nothing - s is a copy
}
```

Range hands you a copy of each element. To modify, index:

```go
for i := range servers {
    servers[i].Active = true
}
```

JavaScript never does this because array elements are references. Go elements are values.

The second one is `append`. A slice header - address, length, capacity - is itself passed by
value, same as everything else. Writing *through* it (`servers[i].Active = true`) reaches the
real backing array, because the address was already on the header you were handed. But `append`
doesn't write through anything - it builds a **new** header, possibly pointing at a new array, and
returns it. Assign that return value to nothing, and the caller never sees it:

```go
func addOne(servers []Server) {
    servers = append(servers, next)   // reassigns the local copy of the header only
}
```

```go
servers = append(servers, next)   // the caller's own reassignment - this one counts
```

`AddServer` below is built around exactly this - it returns the grown slice rather than trying to
grow the one it was handed, because trying the second way is not possible.

## Run

```sh
go test ./collections/
```
