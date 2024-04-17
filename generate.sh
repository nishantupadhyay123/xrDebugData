#!/bin/bash
protoc  -I src/ --go_out=src/ src/xrbinarypb/data.proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    src/xrbinarysrvpb/serv.proto