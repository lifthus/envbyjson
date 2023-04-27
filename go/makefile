.DEFAULT_GOAL := build

test:
	go test -v -cover -coverprofile=testdata/coverage.out
	go tool cover -html=testdata/coverage.out
.PHONY:test