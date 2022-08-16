# GRPC Playground

## Preconditions

- brew install protoc
- go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
- export PATH="$PATH:$(go env GOPATH)/bin"

## Generating python code

- python3 -m venv .venv
- source .venv/bin/activate
- pip install grpcio-tools
- python -m grpc_tools.protoc -I=../ --python_out=. --grpc_python_out=. ../ProductCatalogue.proto

## Go client code

- cd product-catalogue-client
- go mod tidy
- protoc -I .. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ../ProductCatalogue.proto
- go run main.go

## Generate proto docs

docker run --rm \
  -v $(pwd)/doc:/out \
  -v $(pwd):/protos \
  pseudomuto/protoc-gen-doc
