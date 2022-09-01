.PHONY: all build run gotool clean help

.PHONY: all
all: build runpro

.PHONY: build
build:
	@go build -o apiserver ./cmd/apiserver.go

.PHONY: buildpro
buildpro:
	@export GOOS=linux && export GOARCH=amd64 && go build -o apiserver ./cmd/apiserver.go

.PHONY: rundev
rundev:
	@DQENV=dev go run ./cmd/apiserver.go

.PHONY: runpro
runpro:
	@DQENV=pro ./apiserver

.PHONY: help
help:
	@echo "make - 编译生产二进制文件并运行"
	@echo "make build - 本地生成二进制文件"
	@echo "make buildpro - 生产生成二进制文件"
	@echo "make rundev -  本地运行服务"
	@echo "make runpro - 运行线上二进制文件"