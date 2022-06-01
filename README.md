# TodoApp

This is a practice project used to implement things I've learned over time. The todoAPP will have a server and client component.
The server will be written in Golang, while the client will be written in a popular frontend language(TBC).
The server and client with communicate using the gRPC protocol, while the client will also have a REST interface to accept
request from the user. I will also use MySQL to store data persistently.

## Prerequisites

1. Install go 1.17.6 or higher
2. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
3. go install github.com/envoyproxy/protoc-gen-validate

## How to run
```makefile
make proto-gen
make vendor
make run
```

## Testing