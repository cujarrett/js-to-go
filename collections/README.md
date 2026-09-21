# 02 Collections

Slices are arrays. Maps are plain objects. The differences are worth knowing.

| JavaScript | Go |
|---|---|
| `const xs = []` | `var xs []string` (nil, but you can append) |
| `xs.push(x)` | `xs = append(xs, x)` returns a new slice, so reassign |
| `xs.map(f)` | a `for` loop. Go has no map/filter/reduce |
| `const m = {}` | `m := map[string]int{}` |
| `m[k] ?? 0` | `m[k]`, since a missing key already gives the zero value |
| `k in m` | `v, ok := m[k]` |
| `new Set()` | `map[string]bool` |
| `xs.includes(x)` | `slices.Contains(xs, x)`, stdlib since Go 1.21 |
| `Math.min(a, b)` | `min(a, b)`, a builtin since Go 1.21 |
| `a \|\| b \|\| c` | `cmp.Or(a, b, c)`, the first *non-zero* value (`0` and `""` count as zero) |

## Range gives you a copy

```go
for _, s := range servers {
    s.Active = true   // does nothing, s is a copy
}
```

Range hands you a copy of each element. To modify, index:

```go
for i := range servers {
    servers[i].Active = true
}
```

JavaScript never does this because array elements are references. Go elements are values.

The `append` row above is the same idea one level up: it returns a new slice rather than growing
the one you handed it. `AddServer` proves it. The caller only sees the change by reassigning.

## Run

```sh
go test ./collections/
```
