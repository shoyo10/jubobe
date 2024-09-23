pwd=$(shell pwd)

.PHONY: all jubobe unit-test

all: jubobe 

jubobe:
	CONFIG_DIR=${pwd}/configs go run ./main.go jubobe

unit-test:
	go test -count=1 -v ./...
