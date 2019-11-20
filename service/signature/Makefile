compile: clean protoc

protoc:
	protoc -I ../.. ../../proto/Signature.proto --go_out=plugins=grpc:gen
	protoc -I ../.. ../../proto/QR.proto --go_out=plugins=grpc:gen

include ../../base.mk
