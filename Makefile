#!/usr/bin/make -f

export CGO_ENABLED=0

PROJECT=github.com/previousnext/microcron

# Builds the project
build:
	gox -os='linux darwin' -arch='amd64' -output='bin/microcron_{{.OS}}_{{.Arch}}' -ldflags='-extldflags "-static"' $(PROJECT)

# Run all lint checking with exit codes for CI
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

# Run tests with coverage reporting
test:
	go test -cover ./...

IMAGE=previousnext/microcron
VERSION=$(shell git describe --tags --always)

# Releases the project Docker Hub
release:
	docker build -t ${IMAGE}:${VERSION} .
	docker push ${IMAGE}:${VERSION}

.PHONY: build lint test release
