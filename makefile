BINARY_NAME=entproject

build:
	CGO_ENABLED=1 go build -o ${BINARY_NAME} main.go

run:
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

generate:
	go generate ./...
