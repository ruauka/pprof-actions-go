test:
	@go test ./... -coverpkg=./... -cover -coverprofile cover.out
	@echo "-------------------------------------------------------------------------------------"
	@go tool cover -func cover.out
	@echo "-------------------------------------------------------------------------------------"

lint:
	@golangci-lint run

vendor:
	go mod vendor

bench:
	go test -bench=BenchmarkExecute -benchmem -benchtime 5s -count=5

pprof:
	go test -bench=BenchmarkExecute -benchmem -benchtime=5s -count=5 -cpuprofile cpu.out -memprofile mem.out

pprof-mem:
	go tool pprof -http :9000 mem.out

pprof-cpu:
	go tool pprof -http :9010 cpu.out

build:
	docker build --platform=linux/arm64 --tag=ruauka/actions:latest-arm64 .
	#docker build --tag=ruauka/actions:latest .

run:
	docker run -d --rm --name actions -p 8080:8000 actions:latest

remove:
	docker stop actions && docker rmi actions:latest

push:
	docker push ruauka/actions:latest-arm64