#!/usr/bin/env bash

# go get github.com/gogo/protobuf/{proto,protoc-gen-gogo,gogoproto,protoc-gen-gofast,protoc-gen-gogofaster}

protoc \
	--proto_path=$GOPATH/src \
	--proto_path=$GOPATH/src/github.com/gogo/protobuf/protobuf \
		--gogofaster_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:./proto/mail \
		--govalidators_out=gogoimport=true,Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:./proto/mail \
	--proto_path=./proto/mail \
	./proto/mail/*.proto
