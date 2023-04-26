.DEFAULT_GOAL := build

test:
	go test -v -cover -coverprofile=testdata/c.out
	go tool cover -html=testdata/c.out
.PHONY:test