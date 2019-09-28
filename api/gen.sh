#!/usr/bin/env bash

GRPC_GW_PATH=`go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
GRPC_GW_PATH="${GRPC_GW_PATH}/../third_party/googleapis"

# generate the gRPC code
protoc -I. -I${GRPC_GW_PATH} --go_out=plugins=grpc:. \
    device.proto

# generate the JSON interface code
protoc -I. -I${GRPC_GW_PATH} --grpc-gateway_out=logtostderr=true:. \
    device.proto

# generate the swagger definitions
protoc -I. -I${GRPC_GW_PATH} --swagger_out=json_names_for_fields=true:./swagger \
    device.proto

# merge the swagger code into one file
go run swagger/main.go swagger > ../static/swagger/api.swagger.json
