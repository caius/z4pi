GO_FILES = $(shell find . -name \*.go)

.PHONY: default
default: build

.PHONY: build
build: build/z4pi-server

.PHONY: build-native
build-native: build/z4pi-server-arm

build/z4pi-server: $(GO_FILES)
	go build -o build/z4pi-server cmd/server/main.go

build/z4pi-server-arm: $(GO_FILES)
	GOOS=linux GOARCH=arm go build -o build/z4pi-server cmd/server/main.go

.PHONY: clean
clean:
	find build/ -type f -not -name .keep -delete

.PHONY: test
test: $(GO_FILES)
	go test -v ./...

# Useful for debugging sometimes
build/gokbus-arm:
	GOOS=linux GOARCH=arm go build -o build/gokbus-arm github.com/qcasey/gokbus/cmd/gokbus
