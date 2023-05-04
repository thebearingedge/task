include .env
export

MAKEFLAGS += --no-print-directory

.PHONY: test tdd cover run build image

.DEFAULT_GOAL := build

test:
	go test -v ./...

tdd:
	gow -c test -v ./...

cover:
	go test -v ./... -coverprofile .coverage/task.out
	go tool cover -html=.coverage/task.out -o .coverage/task.html

run:
	go run main.go

build:
	go build -o ./.bin/task

image:
	docker build --tag thebearingedge/task .
