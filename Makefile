# Project variables
APP_REST=rest
APP_WORKER=worker
DOCKER_REST_IMAGE=rest
DOCKER_WORKER_IMAGE=worker

# Build Go binaries
build:
	CGO_ENABLED=0 go build -o bin/$(APP_REST) ./cmd/rest
	CGO_ENABLED=0 go build -o bin/$(APP_WORKER) ./cmd/worker

# Build containers with Podman
podman-build:
	podman build -t $(DOCKER_REST_IMAGE) -f Dockerfile.rest .
	podman build -t $(DOCKER_WORKER_IMAGE) -f Dockerfile.worker .

# Run full stack with podman kube
up:
	podman kube play play.yaml

# Tear down stack
down:
	podman kube down play.yaml

# Run REST server directly (for local dev)
run-rest:
	go run ./cmd/rest

# Run Worker directly
run-worker:
	go run ./cmd/worker

# Clean build outputs
clean:
	rm -rf bin/

# Format Go code
fmt:
	go fmt ./...

# Tidy modules
tidy:
	go mod tidy
