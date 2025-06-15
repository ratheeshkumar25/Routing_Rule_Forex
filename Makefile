GOPATH:=$(shell go env GOPATH)

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: run
run:
	@go run cmd/main.go
.PHONY: swagger
swagger:
	@swag fmt -g cmd/main.go
	@swag init -g cmd/main.go