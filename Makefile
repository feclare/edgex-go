.PHONY: build test prepare docker docker_export_client docker_export_distro

EXPORT_CLIENT_VERSION=$(shell cat cmd/export-client/VERSION)
EXPORT_DISTRO_VERSION=$(shell cat cmd/export-distro/VERSION)

build:
	go build -ldflags "-X main.version=$(EXPORT_CLIENT_VERSION)" ./cmd/export-client
	go build -ldflags "-X main.version=$(EXPORT_DISTRO_VERSION)" ./cmd/export-distro
	go build ./cmd/core-data
	go build ./core/metadata
	go build ./core/command

test:
	go test `glide novendor`

prepare:
	glide install

docker_export_client:
	docker build -f docker/Dockerfile.client -t edgexfoundry/docker-export-client .

docker_export_distro:
	docker build -f docker/Dockerfile.distro -t edgexfoundry/docker-export-distro .

docker: docker_export_distro docker_export_client
