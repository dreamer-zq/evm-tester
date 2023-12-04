#!/usr/bin/make -f

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -o build/tester.exe ./cmd
else
	go build -o build/tester ./cmd
endif

build-linux: go.sum
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/tester ./cmd

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify