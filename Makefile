.PHONY: fmt vet lint vuln audit test

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run --verbose

vuln:
	govulncheck ./...

test:
	go test ./... -cover

audit: fmt vet lint vuln test