.PHONY: build test 

EXPORT_CLIENT_VERSION=$(shell cat cmd/export-client/VERSION)
EXPORT_DISTRO_VERSION=$(shell cat cmd/export-distro/VERSION)

build:
	go build -ldflags "-X main.version=$(EXPORT_CLIENT_VERSION)" ./cmd/export-client
	go build -ldflags "-X main.version=$(EXPORT_DISTRO_VERSION)" ./cmd/export-distro
	go build ./core/metadata
	go build ./core/data
	go build ./core/command

test:
	go test `glide novendor`

