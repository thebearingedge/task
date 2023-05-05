include .env
export

MAKEFLAGS += --no-print-directory

.PHONY: test tdd cover run build image

.DEFAULT_GOAL := build

test:
	go test -v ./...

test_only:
	go test -v -run $(t) ./...

tdd:
	gow -c test ./...

tdd_only:
	gow -c test -v -run $(t) ./...

cover:
	go test -v ./... -coverprofile .coverage/task.out
	go tool cover -html=.coverage/task.out -o .coverage/task.html

run:
	go run main.go

watch:
	gow -c run main.go

build:
	go build -o ./.bin/task

image:
	docker build --tag thebearingedge/task .
