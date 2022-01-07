.PHONY: run
run:
	go run main.go -config config.yaml

.PHONY: run-docker
run:
	go run main.go -config config.yaml

.PHONY: test
test:
	go test -v

.PHONY: coverage
coverage:
	go test -v

.PHONY: lint
lint:
	@golint -set_exit_status ${PKG_LIST}