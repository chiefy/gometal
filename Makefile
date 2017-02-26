TEST?=$$(go list ./... | grep -v /vendor/)


default: build


.PHONY: build
build:
	@go build ./main.go

.PHONY: run
run:
	@go run ./main.go


.PHONY: test
test:
	@go test $(TEST)
