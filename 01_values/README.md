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
go test ./01_values/
```

Make all four pass. Then delete an `&` and read the compiler error.
