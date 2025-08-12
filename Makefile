.PHONY: test lint clean

test:
	go test -v ./exercises/...

lint:
	golangci-lint run

clean:
	go clean -testcache
