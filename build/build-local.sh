#!/bin/sh

# This script is for local builds rather than for a build server


export CGO_ENABLED=0

DIRS=`ls -l ./cmd | egrep '^d' | awk '{print $9}'`
for d in ${DIRS}; do
    go build                                                           \
        -ldflags "-X ${PKG}/pkg/version.Str=${VERSION}"                \
        -o ./bin/$(go env GOARCH)/$(go env GOOS)_$(go env GOARCH)/${d}                  \
        ${PKG}/cmd/${d}
done
