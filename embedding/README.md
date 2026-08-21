# 11 Embedding

```js
// JS: inheritance, is-a
class Server extends Metadata {}
new Server().hasLabel("env", "prod")   // inherited, no line written
```

```js
// JS: composition, has-a - the usual advice is to prefer this over extends
class Server {
  constructor(metadata) { this.metadata = metadata }
  hasLabel(k, v) { return this.metadata.hasLabel(k, v) }   // written by hand
}
```

```go
// Go: only has-a exists, but the method still arrives for free
type Server struct {
    Metadata   // no field name - this is the embed
    Active bool
}
server.HasLabel("env", "prod")   // promoted, no method written on Server
```

Go has no `extends`. Embedding gets you the free-method part of inheritance without the is-a part:
`server.HasLabel(...)` works because `HasLabel` is **promoted** onto `Server`, but a function that
takes a `Metadata` still will not accept a `Server` - the type system was never told they are the
same thing, only that one contains the other.

## Why this is everywhere in a controller

```go
type Server struct {
    metav1.ObjectMeta   // Name, Namespace, Labels - all promoted, on every API type
    Spec   ServerSpec
}
```

Every Kubernetes object embeds `ObjectMeta` and `TypeMeta` this way, which is why `server.Name`
works despite no `Name` field ever being declared on `Server` itself.

```go
type Reconciler struct {
    client.Client   // Get, List, Update, ... all promoted
    Scheme *runtime.Scheme
}
```

And it is why a `Reconciler` calls `r.Get(ctx, key, obj)` directly, not `r.Client.Get(...)` - the
embedded `client.Client`'s methods are promoted onto the `Reconciler` itself.

## Run

```sh
go test ./embedding/
```
