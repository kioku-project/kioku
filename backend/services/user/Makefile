GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

.PHONY: proto
proto:
	@protoc --proto_path=../../ --go_out=:../../ --go_opt=module=github.com/kioku-project/kioku services/user/proto/user.proto ../../pkg/proto/common.proto
	@protoc --proto_path=../../ --micro_out=../../ --micro_opt=module=github.com/kioku-project/kioku services/user/proto/user.proto


.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o user *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t backend/services/user:latest --file Dockerfile ../../. 
