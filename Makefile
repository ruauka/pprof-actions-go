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