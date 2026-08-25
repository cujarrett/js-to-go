![logo](./logo.png)

Made this for someone who knows JavaScript and wants to learn Go.

Ten small modules. Stubbed functions, failing tests, and just enough README to point you in the right direction.

Run the test. Write the code. Figure out why it broke.

Start at `values`. Work down.

## Order

| Module | Idea |
|---|---|
| [values](./values/README.md) | pointers, copies, nil |
| [collections](./collections/README.md) | slices, maps, sets |
| [errors](./errors/README.md) | errors as values |
| [types](./types/README.md) | methods, receivers, interfaces |
| [json](./json/README.md) | struct tags, encoding |
| [jsonv2](./jsonv2/README.md) | a real stdlib package, still experimental |
| [slog](./slog/README.md) | structured logging |
| [context](./context/README.md) | cancellation, deadlines, request values |
| [service](./service/README.md) | an HTTP API, stdlib only |
| [cmd](./cmd/README.md) | wiring a binary, graceful shutdown |

## Running

```sh
go test ./values/          # one module
just ci                    # everything
```

## Stuck?

```sh
just diff values      # your answer against the reference, for one module
just solution values  # print the reference answer outright
```

Both read the [`solutions`](https://github.com/cujarrett/js-to-go/tree/solutions) branch without
touching your working tree. Try the test failure first - it's usually enough.
