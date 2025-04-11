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
