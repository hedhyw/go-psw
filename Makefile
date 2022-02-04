GOLANG_CI_LINT_VER:=v1.43.0

lint: bin/golangci-lint
	./bin/golangci-lint run
.PHONY: lint

test:
	go test -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func coverage.out
.PHONY: test.coverage

tidy:
	go mod tidy
.PHONY: vendor

bin/golangci-lint:
	curl \
		-sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s $(GOLANG_CI_LINT_VER)
