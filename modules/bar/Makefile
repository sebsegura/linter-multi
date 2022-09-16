M = $(shell printf "\033[34;1m▶\033[0m")
MODULE=bar

lint: lint-check-deps
	$(info $(M) running linter on $(MODULE) module...)
	@golangci-lint run ./... --config ../../.golangci.yml

lint-check-deps:
	@if [ -z `which golangci-lint` ]; then \
		echo "installing golangci-lint";\
		brew install golangci-lint;\
	fi

format:
	$(info $(M) running gofmt...)
	@go fmt ./...

test:
	@go test -v -race -coverpkg=./... -coverprofile=cov.out -outputdir=./ ./...

coverage:test
	@go tool cover -html=./cov.out -o ./coverage.html
