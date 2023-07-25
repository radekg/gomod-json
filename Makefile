install:
	go build -o $$GOPATH/bin/gomod-json ./cmd/print

.PHONY: test
test:
	go test ./... -count=1 -v
