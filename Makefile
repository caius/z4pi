all:
	@echo "There is no all."

build/gokbus:
	go build -o build/gokbus github.com/qcasey/gokbus/cmd/gokbus

build/gokbus-linux-arm:
	GOOS=linux GOARCH=arm go build -o build/gokbus-linux-arm github.com/qcasey/gokbus/cmd/gokbus

.PHONY: clean
clean:
	find build/ -type f -delete
