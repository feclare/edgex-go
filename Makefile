.PHONY: build test prepare docker docker_export_client docker_export_distro

build:
	go build ./cmd/export-client
	go build ./cmd/export-distro

test:
	go test `glide novendor`

prepare:
	glide install

docker_export_client:
	docker build -f docker/Dockerfile.client -t edgexfoundry/docker-export-client .

docker_export_distro:
	docker build -f docker/Dockerfile.distro -t edgexfoundry/docker-export-distro .

docker: docker_export_distro docker_export_client
