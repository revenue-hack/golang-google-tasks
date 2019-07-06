.PHONY: setup test

setup:
	GO111MODULE=off go get github.com/golang/mock/mockgen

test:
	go test ./...

