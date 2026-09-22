# encoding/json/v2 (module jsonv2) is real but experimental - it only builds under this.
# See jsonv2/README.md.
JSONV2 := "GOEXPERIMENT=jsonv2"

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
    dir=$(just _find {{module}})
    if [ "$dir" = "jsonv2" ]; then
        {{JSONV2}} go test ./jsonv2/
    else
        go test ./"$dir"/
    fi

# Run everything, the way CI does
ci:
    go vet $(go list ./... | grep -v /jsonv2)
    {{JSONV2}} go vet ./jsonv2/
    go test -race $(go list ./... | grep -v /jsonv2)
    {{JSONV2}} go test -race ./jsonv2/

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
        echo "no solutions branch yet - see 'just help-solutions'" >&2
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
        echo "no solutions branch yet - see 'just help-solutions'" >&2
        exit 1
    fi
    # Exercise files only, so a README or test edited since the answer was recorded
    # does not show up as though it were part of your answer.
    git diff solutions -- $(git ls-tree -r --name-only HEAD "$dir/" | grep -v _test.go | grep '\.go$')

# How to record an answer once a module is green
help-solutions:
    @echo "Finish a module on main, leave it uncommitted, then:"
    @echo "  git switch -c solutions      # first time. after that: git switch solutions"
    @echo "  git add <module>/ && git commit -m 'solve <module>'"
    @echo "  git switch main"
    @echo ""
    @echo "Uncommitted work follows you across the switch, so there is nothing to stash."
    @echo "main returns to stubs by itself - you never committed the answer there."
