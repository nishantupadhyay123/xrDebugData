#!/bin/bash
protoc  -I src/ --go_out=src/ src/xrbinarypb/data.proto