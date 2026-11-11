# Finds the one module directory whose name starts with the given argument, e.g. "coll"
# for collections/ - so a short, misspelling-free prefix is enough to type.
_find module:
    #!/usr/bin/env bash
    set -euo pipefail
    dir=$(ls -d "{{module}}"*/ 2>/dev/null | head -1 || true)
    if [ -z "$dir" ]; then
        echo "no module matching '{{module}}'. available:" >&2
        ls -d */ | grep -v '^\.' | sed 's|/$||; s|^|  |' >&2
        exit 1
    fi
    echo "${dir%/}"

# List available recipes
default:
    @just --list

# Run one module's tests, e.g. `just test coll`
test module:
    #!/usr/bin/env bash
    set -euo pipefail
    go test ./"$(just _find {{module}})"/

# Run everything, the way CI does
ci:
    go vet ./...
    go test -race ./...

# Restore exercise files to their unsolved state, e.g. `just reset coll`, or all of them
reset module="":
    #!/usr/bin/env bash
    set -euo pipefail
    scope=""
    if [ -n "{{module}}" ]; then
        scope="$(just _find {{module}})/"
    fi
    # Exercise files only, so tests and READMEs keep any edits made while solving.
    # HEAD is the unsolved state: answers are committed to the solutions branch, never here.
    for f in $(git ls-tree -r --name-only HEAD $scope | grep -v _test.go | grep '\.go$'); do
        git restore --source=HEAD -- "$f"
        echo "reset $f"
    done

# Show the reference answer for one module, e.g. `just solution coll`.
# Reads the solutions branch without touching your working tree.
solution module:
    #!/usr/bin/env bash
    set -euo pipefail
    dir=$(just _find {{module}})
    if ! git rev-parse --verify -q solutions >/dev/null; then
        echo "no solutions branch. Run: git fetch origin solutions:solutions" >&2
        exit 1
    fi
    for f in $(git ls-tree --name-only solutions "$dir/" | grep -v _test.go | grep '\.go$'); do
        git show solutions:"$f"
    done

# Diff your answer against the reference, e.g. `just diff coll`
diff module:
    #!/usr/bin/env bash
    set -euo pipefail
    dir=$(just _find {{module}})
    if ! git rev-parse --verify -q solutions >/dev/null; then
        echo "no solutions branch. Run: git fetch origin solutions:solutions" >&2
        exit 1
    fi
    # Exercise files only, so a README or test edited since the answer was recorded
    # does not show up as though it were part of your answer.
    git diff solutions -- $(git ls-tree -r --name-only HEAD "$dir/" | grep -v _test.go | grep '\.go$')
