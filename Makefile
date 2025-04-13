.PHONY: gen
gen:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto
	@mockery

.PHONY: test
test:
	@mkdir tmp 2>/dev/null || true
	@go test -race -v -coverprofile=tmp/coverage.out -count=1 ./usecase/... 
	@go tool cover -func=tmp/coverage.out
	@go tool cover -html=tmp/coverage.out -o tmp/coverage.html

.PHONY: coverage
coverage:
	@open tmp/coverage.html

.PHONY: docker-build
docker-build:
	@docker build -t dailoi2807/vrs-general-service .

.PHONY: docker-run
docker-run:
	@docker run -p 9001:9001 dailoi2807/vrs-general-service

.PHONY: dev
dev:
	@air main.go
