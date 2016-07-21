#!/bin/sh
#
# Generate test coverage statistics for Go packages.
#
# Works around the fact that `go test -coverprofile` currently does not work
# with multiple packages, see https://code.google.com/p/go/issues/detail?id=6909
#
# Usage: script/coverage [-v]
#
#     -v          Verbose Test Output
#

set -e

workdir=.cover
profile="$workdir/cover.out"
mode=count
verbose=false
showhtml=false
timeout=30s

generate_cover_data() {
    rm -rf "$workdir"
    mkdir "$workdir"

    v=""
    if [ "$verbose" = true ] ; then
        v="-v "
    fi

    for pkg in "$@"; do
        f="$workdir/$(echo $pkg | tr / -).cover"
        go test $v -timeout="$timeout" -covermode="$mode" -coverprofile="$f" "$pkg"
    done

    echo "mode: $mode" >"$profile"
    grep -h -v "^mode:" "$workdir"/*.cover >>"$profile"
}

show_cover_report() {
    go tool cover -${1}="$profile"
}

case "$1" in
"")
    ;;
-v)
    verbose=true ;;
--html)
    showhtml=true ;;
esac

case $2 in
"")
    ;;
--html)
    showhtml=true ;;
esac

generate_cover_data $(go list ./... | grep -v /vendor/)

show_cover_report func

if [ "$showhtml" = true ] ; then
    show_cover_report html
fi
