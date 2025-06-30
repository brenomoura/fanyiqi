install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.7 &&\
	go install github.com/fyne-io/fyne-cross@v1.6.1
	go mod tidy && \
	go mod vendor

run:
	go run cmd/gui/main.go

lint:
	golangci-lint run ./...

build-linux:
	fyne-cross linux -icon ui/assets/fanyiqi_icon.png

build-windows:
	fyne-cross windows -app-id fanyiqi.windows -icon ui/assets/fanyiqi_icon.png

build-macos:
	fyne-cross darwin -app-id fanyiqi.macos -icon ui/assets/fanyiqi_icon.png