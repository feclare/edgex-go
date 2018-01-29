.PHONY: build test 

EXPORT_CLIENT_VERSION=$(shell cat cmd/export-client/VERSION)
EXPORT_DISTRO_VERSION=$(shell cat cmd/export-distro/VERSION)

build:
	go build -ldflags "-X main.version=$(EXPORT_CLIENT_VERSION)" ./cmd/export-client
	go build -ldflags "-X main.version=$(EXPORT_DISTRO_VERSION)" ./cmd/export-distro

test:
	go test `glide novendor`

