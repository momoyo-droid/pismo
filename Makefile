.PHONY: build fmt vet lint vuln audit test

build:
	docker compose down
	docker compose build --no-cache
	docker compose up -d

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