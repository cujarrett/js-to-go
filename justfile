# Git tag marking the commit where every function is still a TODO. `just reset` restores
# from here, so it must never be moved onto a commit that contains answers.
# Re-point it with: git tag -f stubs <sha>
STUBS := "stubs"

# List available recipes
default:
    @just --list

# Run one module's tests, e.g. `just test 02`
test module:
    go test ./{{module}}_*/

# Run everything, the way CI does
ci:
    go vet ./...
    go test -race ./...

# Restore one module's exercise file to its unsolved state, e.g. `just reset 02`.
# Only the file you edit is replaced - tests and READMEs are left alone.
reset module:
    #!/usr/bin/env bash
    set -euo pipefail
    # Accept 3, 03 or 03_errors - a single digit is the obvious thing to type.
    m="{{module}}"
    if [ ${#m} -eq 1 ]; then m="0$m"; fi
    # || true so a no-match does not trip pipefail before the message below runs.
    dir=$(ls -d "$m"*/ 2>/dev/null | head -1 || true)
    if [ -z "$dir" ]; then
        echo "no module matching '{{module}}'. available:" >&2
        ls -d 0*_*/ | sed 's|/$||; s|^|  |' >&2
        exit 1
    fi
    dir=${dir%/}
    for f in $(git ls-tree --name-only {{STUBS}} "$dir/" | grep -v _test.go | grep '\.go$'); do
        git show {{STUBS}}:"$f" > "$f"
        echo "reset $f"
    done

# Restore every module. Wipes all your answers.
reset-all:
    #!/usr/bin/env bash
    set -euo pipefail
    for f in $(git ls-tree -r --name-only {{STUBS}} | grep '^0[0-9]_' | grep -v _test.go | grep '\.go$'); do
        git show {{STUBS}}:"$f" > "$f"
    done
    echo "reset all six modules"
