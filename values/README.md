# 01 Values

Go makes you say whether you mean the thing or a copy of the thing. JavaScript never asks.

The same function in each language.

```js
// JS: objects are always the real thing
const activate = (s) => {
  s.active = true   // caller sees it
}
```

```go
// Go: a pointer, so the same object
func Activate(s *Server) {
    s.Active = true   // caller sees it
}
```

`*Server` in a parameter list means "a pointer to a Server", the address of one rather than a
Server itself. It is how you ask for what JS handed you for free above.

Go's second option has no JS equivalent. You can ask for a copy instead.

```go
// Go: no pointer, so a copy
func activateCopy(s Server) {
    s.Active = true   // caller does not see it, s is a copy
}
```

There is no exercise for that one. Nothing a copy does is visible from outside, so no test could
tell a correct version from an empty one. `SetLast` below is the same lesson, made testable.

Calling them is where it shows:

```go
srv := Server{Name: "web-1"}
Activate(&srv)       // &srv is "the address of srv"
srv.Active           // true

srv2 := Server{Name: "web-2"}
activateCopy(srv2)   // no &, so Go passes a copy
srv2.Active          // false
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

You write `&` often and `*` rarely, because a field access dereferences for you. A `*string` has
no fields to reach through, so `*n` is the only way in.

Arrays work the same as `Server`. `[3]string` is data, not an address. `SetLast`, one of the
functions you write below, has to return the array because there is nothing else it could do.

`&srv` takes the address of a variable you already have. `new` is for when you have no variable to
point at: it allocates a value and returns its **address**.

```go
s := new(string)   // a fresh "", s is a *string
p := new(5)        // a fresh 5, p is a *int. Go 1.26 and later
```

Hand `new` a variable and it copies it, the same as passing it to any other function. Writing
through the pointer leaves the original alone:

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

One function at a time is enough. Run the test, read the top failure, fix that function, run it
again. Stop whenever. The functions are in the order this README explains them.

When it is all green, delete an `&` and read the compiler error.
