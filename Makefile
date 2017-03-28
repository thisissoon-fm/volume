#
#  Make binaries and run test suite
#

GOOUTDIR 		?= .
GOOS 			?=
GOARCH 			?=
TIME 			?= $(shell date +%s)
VERSION 		?= $(shell git rev-parse HEAD)
CGO_ENABLED 	?= 0

.PHONY: linux darwin

# All Targets
all: linux darwin

# Build Linux
linux%: GOOS = linux
linux%: CGO_ENABLED = 1
linux: linux64

# Build Darwin
darwin%: GOOS = darwin
darwin: darwin64

# 64bit Archetecture
%64: GOARCH = amd64

# Common Build Target
linux64 darwin64:
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	CGO_ENABLED=$(CGO_ENABLED) \
	go build -v \
		-ldflags "-X volume/build.timestamp=$(TIME) -X volume/build.version=$(VERSION) -X volume/build.arch=$(GOARCH) -X volume/build.os=$(GOOS)" \
		-o "$(GOOUTDIR)/sfmvolume.$(GOOS)-$(GOARCH)"
