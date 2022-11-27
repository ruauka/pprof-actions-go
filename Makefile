test:
	@go test -cover ./... -coverprofile cover.out
	@echo "-------------------------------------------------------------------------------------"
	@go tool cover -func cover.out
	@echo "-------------------------------------------------------------------------------------"

lint:
	@golangci-lint run

vendor:
	go mod vendor

dockerup:
	docker-compose up -d --build

dockerstop:
	docker-compose stop

build:
	docker build --tag=actions:latest .

run:
	docker run -d --rm --name actions -p 8080:8000 actions:latest

remove:
	docker stop actions && docker rm -f actions && docker rmi actions:latest