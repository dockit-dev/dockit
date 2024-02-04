build:
	docker build -t dockit-cli .

container: build
	docker run -it dockit-cli

test: build
	docker run dockit-cli go test ./...
