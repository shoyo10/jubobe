pwd=$(shell pwd)

.PHONY: all jubobe unit-test

all: jubobe 

jubobe:
	CONFIG_DIR=${pwd}/configs go run ./main.go jubobe

unit-test:
	go test -count=1 -v ./...

pgmigrate:
	CONFIG_DIR=${pwd}/configs go run ./main.go pgmigration migrate

pgrollback:
	CONFIG_DIR=${pwd}/configs go run ./main.go pgmigration rollback

local-deploy-up:
	docker-compose -f deployment/docker-compose.yml up -d

local-deploy-down:
	docker-compose -f deployment/docker-compose.yml down
