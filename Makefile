.PHONY: build test prepare docker docker_export_client docker_export_distro

GO=CGO_ENABLED=0 go
GOCGO=CGO_ENABLED=1 go

EXPORT_CLIENT_VERSION=$(shell cat cmd/export-client/VERSION)
EXPORT_DISTRO_VERSION=$(shell cat cmd/export-distro/VERSION)
CORE_DATA_VERSION=$(shell cat cmd/core-data/VERSION)
CORE_METADATA_VERSION=$(shell cat cmd/core-metadata/VERSION)
CORE_COMMAND_VERSION=$(shell cat cmd/core-command/VERSION)

MICROSERVICES=cmd/core-data/core-data cmd/core-metadata/core-metadata \
	cmd/core-command/core-command cmd/export-client/export-client \
	cmd/export-distro/export-distro
.PHONY: $(MICROSERVICES)


build: $(MICROSERVICES)

cmd/core-data/core-data:
	$(GOCGO) build -ldflags "-X main.version=$(CORE_DATA_VERSION)" -o cmd/core-data/core-data ./cmd/core-data

cmd/core-metadata/core-metadata:
	$(GO) build -ldflags "-X main.version=$(CORE_METADATA_VERSION)" -o cmd/core-metadata/core-metadata ./cmd/core-metadata

cmd/core-command/core-command:
	$(GO) build -ldflags "-X main.version=$(CORE_COMMAND_VERSION)" -o cmd/core-command/core-command ./cmd/core-command

cmd/export-client/export-client:
	$(GO) build -ldflags "-X main.version=$(EXPORT_CLIENT_VERSION)" -o cmd/export-client/export-client ./cmd/export-client

cmd/export-distro/export-distro:
	$(GOCGO) build -ldflags "-X main.version=$(EXPORT_DISTRO_VERSION)" -o cmd/export-distro/export-distro ./cmd/export-distro

test:
	go test `glide novendor`

prepare:
	glide install

docker_export_client:
	docker build -f docker/Dockerfile.client -t edgexfoundry/docker-export-client .

docker_export_distro:
	docker build -f docker/Dockerfile.distro -t edgexfoundry/docker-export-distro .

docker: docker_export_distro docker_export_client
