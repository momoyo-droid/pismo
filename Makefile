.PHONY: swagger build fmt vet lint vuln audit test

swagger:
	swag init -g api/cmd/app/main.go

build:
	docker compose down
	docker compose build --no-cache
	docker compose up

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