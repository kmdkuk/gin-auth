.PHONY: start
start:
	docker-compose up -d

.PHONY: build
build:
	go build -o bin/gin-auth
