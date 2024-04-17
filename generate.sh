#!/bin/bash
protoc  --go_out=.  --go_opt=paths=source_relative src/xrbinarypb/data.proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    src/xrbinarysrvpb/serv.proto