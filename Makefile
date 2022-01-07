.PHONY: run
run:
	go run main.go -config config.yaml

.PHONY: build-docker
build-docker:
	docker build -t pex .

.PHONY: run-docker
run-docker:
	docker run -dp 8080:8080 pex

.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test ./... -coverprofile coverage.out

.PHONY: lint
lint:
	@golint