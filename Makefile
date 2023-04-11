.PHONY: all clean build-windows build-linux

all: build-windows build-linux

build-windows:
	set GOOS=windows&& set GOARCH=amd64&& go build -o convjson.exe

build-linux:
	set CGO_ENABLED=0&& set GOOS=linux&& set GOARCH=amd64&& go build -o convjson
