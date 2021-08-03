CWD := $(shell pwd)

.PHONY: build 
build:
	docker build -t motoyukikimura/go_linux .

.PHONY: up
up:
	docker run --name go_linux -d -v $(CWD):/root/go/src/github.com/mox692/gooo/asm -it motoyukikimura/go_linux bash

 .PHONY: down
down:
	docker stop go_linux && docker rm -f go_linux 