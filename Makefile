.PHONY: start
start:
	docker-compose build
	docker-compose up -d

.PHONY: logs-api
logs-api:
	docker-compose logs -f api

.PHONY: build
build:
	go build -o bin/gin-auth

.PHONY: user-create
user-create:
	curl -X 'POST' 'http://localhost:3000/users' -H 'accept: */*' -H 'Content-Type: application/json' -d '{"user_id": "string","password": "string"}'

.PHONY: user-login
user-login:
	curl -v -X 'POST' 'http://localhost:3000/login' -H 'accept: */*' -H 'Content-Type: application/json' -d '{"user_id": "string", "password": "string" }'

MYSESSION=test
.PHONY: user-current
user-current:
	curl -X 'GET' 'http://localhost:3000/users' -H 'accept: */*' -H 'Cookie: MYSESSION=${MYSESSION}'

.PHONY: test
test: 
	go test -race -timeout 30m -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: install-go-tools
install-go-tools:
	cat tools.go | awk -F'"' '/_/ {print $$2}' | xargs -tI {} go install {}
