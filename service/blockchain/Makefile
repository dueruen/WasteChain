compile: clean protoc

protoc:
	protoc -I ../.. ../../proto/Blockchain.proto --go_out=plugins=grpc:gen
	protoc -I ../.. ../../proto/Signature.proto --go_out=plugins=grpc:gen

include ../../base.mk
