.DEFAULT_GOAL := run

build:
	go build

run:
	./stdoutinator | go run main.go