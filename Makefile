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
	curl -X 'POST' 'http://localhost:3000/users' -H 'accept: */*' -H 'Content-Type: application/json' -d '{"id": "string","password": "string"}'

.PHONY: user-login
user-login:
	curl -v -X 'POST' 'http://localhost:3000/login' -H 'accept: */*' -H 'Content-Type: application/json' -d '{"id": "string", "password": "string" }'

MYSESSION=test
.PHONY: user-current
user-current:
	curl -X 'GET' 'http://localhost:3000/users' -H 'accept: */*' -H 'Cookie: MYSESSION=${MYSESSION}'
