#!/usr/bin/make -f

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -o build/tester.exe ./cmd
else
	go build -o build/tester ./cmd
endif

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify