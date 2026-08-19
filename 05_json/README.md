# 05 JSON and tags

In JavaScript, the field name in code is the field name on the wire. In Go they are separate, and
the struct tag is the translation.

```js
// JS: the object is the shape, so there is nothing to declare
const config = { sourceSecret: "demo1-tls" }
JSON.stringify(config)   // {"sourceSecret":"demo1-tls"}
```

```go
// Go: the shape is a type, and the tag is the wire name
type Config struct {
    SourceSecret string `json:"sourceSecret"`
}

config := Config{SourceSecret: "demo1-tls"}
json.Marshal(config)   // {"sourceSecret":"demo1-tls"}
```

Two rules force this apart:

- a Go field must start with a **capital letter** to be visible outside its package
- Kubernetes and most JSON APIs use **lowercase**

So every serialisable Go struct carries tags. Lowercase fields are invisible to the encoder
entirely, which is a useful way to keep something out of the wire format.

## omitempty

`json:"replicas,omitempty"` drops the key when the value is the zero value. Good for optional
scalars, usually wrong for lists: a missing key and an empty list mean different things to a
reader, and `targets: []` is a real statement.

This is the same decision as `copies int32` in secret-mirror-controller, where `0` is a genuine
answer and had to stay visible.

## Run

```sh
go test ./05_json/
```

Then rename a tag and watch a test fail. That is how a CRD field gets renamed.
