build:
	docker build -t dockit-cli .

container: build
	docker run -it dockit-cli

test: build
	docker run -v $(PWD):/coverage dockit-cli go test -covermode=atomic -coverprofile=/coverage/coverage.out ./...
