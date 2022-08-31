.PHONY: all build run gotool clean help

.PHONY: build
build:
	@go build -o bin/apiserver ./cmd/apiserver.go

.PHONY: buildpro
buildpro:
	@export GOOS=linux && export GOARCH=amd64 && go build -o bin/apiserver ./cmd/apiserver.go

.PHONY: rundev
run:
	@DQENV=dev go run ./cmd/apiserver.go
