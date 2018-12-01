# This repo's root import path (under GOPATH).
PKG := gitlab.com/leononame/advent-of-code-2018

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)

local:
	PKG=${PKG} VERSION=${VERSION} ./build/build-local.sh

clean:
	rm -rf .go bin