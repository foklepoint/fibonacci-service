default: build-docker-static-binary-image

build-image:
	@docker build --file Dockerfile.build --tag fibonacci-service:build .

bin/main: build-image
	@docker run --rm \
						--volume $(shell pwd):/app fibonacci-service:build \
						/bin/bash -c 'CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o bin/main .'
build-docker-static-binary-image: bin/main
	@docker build --file Dockerfile.static --tag fibonacci-service:latest .
clean:
	@docker volume rm `docker volume ls -q`
