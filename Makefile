default: test

.PHONY: test
test:
	go test -v -race -cover ./...
