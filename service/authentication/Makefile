compile: clean protoc

protoc:
	protoc -I ../.. ../../proto/Authentication.proto --go_out=plugins=grpc:gen

build: compile
	cd cmd && GOOS=linux go build main.go
	docker build -t authentication-service .

include ../../base.mk
