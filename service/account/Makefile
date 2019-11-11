compile: clean protoc

protoc:
	protoc -I ../.. ../../proto/Account.proto --go_out=plugins=grpc:gen

include ../../base.mk
