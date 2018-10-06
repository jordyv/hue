.PHONY: all build build-arm

all: build build-arm

build-arm:
	GOARCH=arm GOOS=linux go build -o bin/hue_arm main.go

build:
	go build -o bin/hue main.go
