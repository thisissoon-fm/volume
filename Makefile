#
#  Make binaries and run test suite
#

GOOUTDIR 		?= .
GOOS 			?=
GOARCH 			?=
TIME 			?= $(shell date +%s)
VERSION 		?= $(shell git rev-parse HEAD)
CGO_ENABLED 	?= 0
CGO_CFLAGS 		?= ""
CGO_LDFLAGS 	?= ""
LDFLAGS 		?= ""

.PHONY: linux darwin

# Build RPXC container
rpxc:
	docker build -t registry.soon.build/sfm/volume:rpxc -f Dockerfile.rpxc .
	docker run \
		--rm \
		-it \
		-v `pwd`:/go/src/volume \
		-e TIME=$(TIME) \
		-e VERSION=$(VERSION) \
		registry.soon.build/sfm/volume:rpxc

# All Targets
all: linux darwin

# Build Linux
linux%: GOOS = linux
linux: linux64

# Build Darwin
darwin%: GOOS = darwin
darwin: darwin64

# 64bit Archetecture
%64: GOARCH = amd64

# ARM Archetecture
arm: GOARCH = arm
arm: CGO_ENABLED=1

# Common Build Target
arm linux64 darwin64:
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	CGO_ENABLED=$(CGO_ENABLED) \
	CGO_LDFLAGS="$(CGO_LDFLAGS)" \
	CGO_CFLAGS="$(CGO_CFLAGS)" \
	go build -v \
		-ldflags "-X volume/build.timestamp=$(TIME) -X volume/build.version=$(VERSION) -X volume/build.arch=$(GOARCH) -X volume/build.os=$(GOOS)" \
		-o "$(GOOUTDIR)/sfmvolume.$(GOOS)-$(GOARCH)"
