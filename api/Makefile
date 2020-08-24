all: build

export GO111MODULE	?= on

dep:
	GO111MODULE=on go mod download

tidy:
	go mod tidy
	gofmt -s -w ./**/*.go
	goimports -w ./**/*.go

lint:
	golangci-lint run

build: dep
	GO111MODULE=on CGO_ENABLED=0 go build -ldflags="-s -w" \
	-o build/uinvest .

run: build
	./build/uinvest

generate:
	go run github.com/99designs/gqlgen