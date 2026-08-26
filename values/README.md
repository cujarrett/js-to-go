# 01 Values

Go makes you say whether you mean the thing or a copy of the thing. JavaScript never asks.

Same function, same two parameters, in each language.

```js
// JS: objects are always the real thing
const rename = (s, name) => {
  s.name = name   // caller sees it
}
```

```go
// Go: a pointer, so the same object
func Rename(s *Server, name string) {
    s.Name = name   // caller sees it
}
```

`*Server` in a parameter list means "a pointer to a Server" - the address of one, rather than a
Server itself. It is how you ask for what JS handed you for free above.

Go's second option has no JS equivalent, and that is the module: you can ask for a copy instead,
and JS cannot.

```go
// Go: no pointer, so a copy
func RenameCopy(s Server, name string) {
    s.Name = name   // caller does not see it, s is a copy
}
```

Calling them is where it shows:

```go
srv := Server{Name: "old"}
Rename(&srv, "renamed")    // &srv is "the address of srv"
srv.Name                   // renamed

srv2 := Server{Name: "old"}
RenameCopy(srv2, "renamed")   // no &, so Go passes a copy
srv2.Name                     // old
```

`&` makes a pointer, `*` in a type means one. You write `&` often. You almost never write `*` to
read a field, because Go dereferences for you: `s.Name` works whether `s` is a `Server` or a
`*Server`.

Arrays work the same as `Server` - `[3]string` is data, not an address. `SetFirst` proves it: it
returns the changed array, because there is nothing else it could do.

`&srv` is the only way to get a pointer to a variable you already have. Go 1.26 added a second way
for a value you don't have a variable for yet: `new(5)` returns a `*int` pointing at a fresh `5`.
It is not `&` in disguise, though - `new(x)` makes a **copy** of `x` and hands back a pointer to
that copy, the same as passing `x` to any other function. Mutate through the pointer, and the
original `x` does not change:

```go
x := 5
p := new(x)
*p = 99
x   // still 5 - p points at a copy, not at x
```

## nil

A pointer can be nil, which is how an API says "not set". That is different from `""`.

```js
// JS has two of these already
{ note: undefined }   // not set
{ note: "" }          // set to empty
```

```go
Note *string   // nil = not set, &"" = set to empty
```

Reading `*s.Note` when it is nil panics. Check first.

## Run

```sh
go test ./values/
```

Make every test pass. Then delete an `&` and read the compiler error.
