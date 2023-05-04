set dotenv-load

[private]
default:
  @just --list

test:
  go test -v ./...

tdd:
  gow -c test -v ./...

cover:
  go test -v ./... -coverprofile .coverage/task.out
  go tool cover -html=.coverage/task.out -o .coverage/task.html

build:
  go build -o ./.bin/task
