.PHONY: vet test

test:
	go test ./...

test/coverage:
	go test -cover ./...

vet:
	go vet ./...
