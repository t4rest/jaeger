#!/usr/bin/env bash

protoc \
	--proto_path=$GOPATH/src \
	--proto_path=$GOPATH/src/github.com/gogo/protobuf/protobuf \
		--gogofaster_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:./proto/rules \
		--govalidators_out=gogoimport=true,Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:./proto/rules \
	--proto_path=./proto/rules \
	./proto/rules/*.proto


protoc \
	--proto_path=$GOPATH/src \
	--proto_path=$GOPATH/src/github.com/gogo/protobuf/protobuf \
		--gogofaster_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:./proto/subscription \
		--govalidators_out=gogoimport=true,Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:./proto/subscription \
	--proto_path=./proto/subscription \
	./proto/subscription/*.proto

protoc \
	--proto_path=$GOPATH/src \
	--proto_path=$GOPATH/src/github.com/gogo/protobuf/protobuf \
		--gogofaster_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:./proto/group_changes \
		--govalidators_out=gogoimport=true,Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:./proto/group_changes \
	--proto_path=./proto/group_changes \
	./proto/group_changes/*.proto