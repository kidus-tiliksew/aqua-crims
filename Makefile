.PHONY: all core notification

all: core notification

core:
	go run cmd/core/main.go

notification:
	go run cmd/notification/main.go

cli:
	go run cmd/cli/main.go