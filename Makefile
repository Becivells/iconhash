# This how we want to name the binary output
#
# use checkmake linter https://github.com/mrtazz/checkmake
# $ checkmake Makefile

BINARY=iconhash
GITREPO=Becivells/iconhash
GOPATH ?= $(shell go env GOPATH)
# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
PATH := ${GOPATH}/bin:$(PATH)
GCFLAGS=-gcflags "all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}"

VERSION_TAG := $(shell git describe --tags --always)
VERSION_VERSION := $(shell git log --date=iso --pretty=format:"%cd" -1) $(VERSION_TAG)
VERSION_COMPILE := $(shell date +"%F %T %z")
VERSION_BRANCH  := $(shell git rev-parse --abbrev-ref HEAD)
VERSION_GIT_DIRTY := $(shell git diff --no-ext-diff 2>/dev/null | wc -l|sed 's/^[ \t]*//g')
VERSION_DEV_PATH:= $(shell pwd)
LDFLAGS=-ldflags=" -s -w -X 'main.VERSION_TAG=$(VERSION_TAG)' -X 'main.version=$(VERSION_VERSION)' -X 'main.date=$(VERSION_COMPILE)' -X 'main.Branch=$(VERSION_BRANCH)' -X 'main.GitDirty=$(VERSION_GIT_DIRTY)'"

# These are the values we want to pass for VERSION  and BUILD
BUILD_TIME=`date +%Y%m%d%H%M`
COMMIT_VERSION=`git rev-parse HEAD`
PLDCNAME=pldc-$(VERSION_TAG)-$(VERSION_BRANCH)-$(VERSION_GIT_DIRTY)-$(BUILD_TIME)
# colors compatible setting
COLOR_ENABLE=$(shell tput colors > /dev/null; echo $$?)
ifeq "$(COLOR_ENABLE)" "0"
CRED=$(shell echo "\033[91m")
CGREEN=$(shell echo "\033[92m")
CYELLOW=$(shell echo "\033[93m")
CEND=$(shell echo "\033[0m")
endif

.PHONY: all
all: | fmt build

.PHONY: go_version_check
GO_VERSION_MIN=1.11
# Parse out the x.y or x.y.z version and output a single value x*10000+y*100+z (e.g., 1.9 is 10900)
# that allows the three components to be checked in a single comparison.
VER_TO_INT:=awk '{split(substr($$0, match ($$0, /[0-9\.]+/)), a, "."); print a[1]*10000+a[2]*100+a[3]}'
go_version_check:
	@echo "$(CGREEN)Go version check ...$(CEND)"
	@if test $(shell go version | $(VER_TO_INT) ) -lt \
  	$(shell echo "$(GO_VERSION_MIN)" | $(VER_TO_INT)); \
  	then printf "go version $(GO_VERSION_MIN)+ required, found: "; go version; exit 1; \
		else echo "go version check pass";	fi

# Code format
.PHONY: fmt
fmt: go_version_check
	@echo "$(CGREEN)Run gofmt on all source files ...$(CEND)"
	go generate
	@echo "gofmt -l -s -w ..."
	@ret=0 && for d in $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		gofmt -l -s -w $$d/*.go || ret=$$? ; \
	done ; exit $$ret

# Builds the project
build: fmt
	@echo "$(CGREEN)Building ...$(CEND)"
	@rm -rf bin
	@mkdir -p bin
	CGO_ENABLED=0  go build ${GCFLAGS} ${LDFLAGS} -o bin/${BINARY}
	@echo "build Success!"

.PHONY: release
release: build
	@echo "$(CGREEN)Cross platform building for release ...$(CEND)"
	@mkdir -p release
	@for GOOS in darwin linux windows freebsd netbsd openbsd plan9 solaris; do \
		for GOARCH in amd64 386 arm; do \
			for d in $$(go list -f '{{if (eq .Name "main")}}{{.ImportPath}}{{end}}' ./...); do \
				b=$$(basename $${d}) ; \
				if [ "$${GOOS}" = 'windows' ]; then\
				echo "Building $${b}.$${GOOS}-$${GOARCH}.exe ..."; \
				GOOS=$${GOOS} GOARCH=$${GOARCH} go build ${GCFLAGS} ${LDFLAGS} -v -o release/$${b}.$${GOOS}-$${GOARCH}.exe $$d 2>/dev/null ; \
				cd release &&shasum $${b}.$${GOOS}-$${GOARCH}.exe>$${b}.$${GOOS}-$${GOARCH}.exe.shasum && tar -zcf $${b}.$${GOOS}-$${GOARCH}.tar.gz $${b}.$${GOOS}-$${GOARCH}.exe $${b}.$${GOOS}-$${GOARCH}.exe.shasum;\
				cd ../; \
				else \
				echo "Building $${b}.$${GOOS}-$${GOARCH} ..."; \
				GOOS=$${GOOS} GOARCH=$${GOARCH} go build ${GCFLAGS} ${LDFLAGS} -v -o release/$${b}.$${GOOS}-$${GOARCH} $$d 2>/dev/null ; \
				cd release &&shasum $${b}.$${GOOS}-$${GOARCH}>$${b}.$${GOOS}-$${GOARCH}.shasum && tar -zcf $${b}.$${GOOS}-$${GOARCH}.tar.gz $${b}.$${GOOS}-$${GOARCH} $${b}.$${GOOS}-$${GOARCH}.shasum; \
				cd ../; \
				fi \
			done ; \
		done ;\
	done
	@find ./release/ -type f -a -size 0 -exec rm {} \;

.PHONY: test-cli
test-cli: build
	@echo "$(CGREEN)Run all cli test cases ...$(CEND)"
	bats ./test
	@echo "test-cli Success!"

.PHONY: clean
clean:
	@echo "$(CGREEN)clean all ...$(CEND)"
	@find ./ -type f -a -name iconhash -exec rm {} \;
	@find ./ -name .DS_Store -a -type f -exec rm -f {} \;
	@find ./release/  -type f -exec rm -f {} \;
	@find ./tmp/   -type f -exec rm -f {} \;
	@find ./bin/   -type f -exec rm -f {} \;