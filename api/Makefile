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
	-o build/uinvest cmd/main.go

run: build
	./build/uinvest

test:
	go test -v ./... -cover -vet -all -coverprofile=coverage.out

cover:
	go tool cover -html=coverage.out

generate:
	go run github.com/99designs/gqlgen