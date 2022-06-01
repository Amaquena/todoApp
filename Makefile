PROTO_PATH=./proto

.PHONY: vendor
vendor:
	go mod tidy -compat=1.17
	go mod vendor

.PHONY: run
run:
	go run -mod=vendor ./cmd/...

.PHONY: proto-gen
proto-gen:
	protoc \
		-I ${PROTO_PATH} \
		-I ${GOPATH}/bin/protoc-gen-validate \
		--go_out=paths=source_relative:${PROTO_PATH} \
		--validate_out=paths=source_relative,lang=go:${PROTO_PATH}  \
		--go-grpc_out=${PROTO_PATH} --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		${PROTO_PATH}/todo.proto

.PHONY: proto-validate
proto-validate:
	protoc \
		-I ${PROTO_PATH} \
        -I ${PROTO_PATH} \
        --go_out=paths=source_relative:${PROTO_PATH} \
        --validate_out=paths=source_relative,lang=go:${PROTO_PATH} \
        ${PROTO_PATH}/todo.proto