![logo](./logo.png)

Made this for someone who knows JavaScript and wants to learn Go.

Ten small modules. Stubbed functions, failing tests, and just enough README to point you in the right direction.

Run the test. Write the code. Figure out why it broke.

Start at `values`. Work down.

## Order

| Module | Idea |
|---|---|
| [values](./values) | pointers, copies, nil |
| [collections](./collections) | slices, maps, sets |
| [errors](./errors) | errors as values |
| [types](./types) | methods, receivers, interfaces |
| [json](./json) | struct tags, encoding |
| [jsonv2](./jsonv2) | a real stdlib package, still experimental |
| [slog](./slog) | structured logging |
| [context](./context) | cancellation, deadlines, request values |
| [service](./service) | an HTTP API, stdlib only |
| [cmd](./cmd) | wiring a binary, graceful shutdown |

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
