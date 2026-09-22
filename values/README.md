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

`*Server` in a parameter list means "a pointer to a Server", the address of one rather than a
Server itself. It is how you ask for what JS handed you for free above.

Go's second option has no JS equivalent. You can ask for a copy instead.

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

`&` makes a pointer, `*` in a type means one. The same `*` in front of a value does the opposite
job, so read it by position:

```go
s := Server{Name: "old"}
p := &s              // & in front of a value: the address of s. p has type *Server
p.Name = "web-1"     // no * needed, Go dereferences a field access for you
s.Name               // web-1, because p and s are the same Server

n := &s.Name         // a field has an address like anything else. n has type *string
*n = "web-2"         // * in front of a value: the string n points at
s.Name               // web-2
```

- `*Server` in a type: a pointer to a `Server`.
- `&s` in front of a value: the address of `s`.
- `*n` in front of a value: the thing `n` points at.

You write `&` often. You almost never write `*` to read a field, because Go dereferences for you:
`s.Name` works whether `s` is a `Server` or a `*Server`. A `*string` has no fields to reach
through, so `*n` is the only way to get at the value.

Arrays work the same as `Server`. `[3]string` is data, not an address. `SetLast` proves it: it
returns the changed array, because there is nothing else it could do.

`&srv` is the only way to get a pointer to a variable you already have. Go 1.26 added a second way
for a value you don't have a variable for yet: `new(5)` returns a `*int` pointing at a fresh `5`.
`new(x)` makes a **copy** of `x` and hands back a pointer to that copy, the same as passing `x`
to any other function. Mutate through the pointer, and the original `x` does not change:

```go
x := 5
p := new(x)
*p = 99
x   // still 5, p points at a copy of x
```

## nil

A pointer can be nil, which is how an API says "not set". That is different from `""`.

Go has no `undefined`. A declared variable is its zero value immediately, and the zero value of a
pointer is nil, so `var n *string` is already nil with no init step to write. `n := nil` does not
compile, because `nil` on its own has no type for Go to infer.

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
