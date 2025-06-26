install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.7 &&\
	go mod tidy && \
	go mod vendor

# API_URL ?= http://localhost:8000/api/v1

run:
	go run cmd/gui/main.go

lint:
	golangci-lint run ./...