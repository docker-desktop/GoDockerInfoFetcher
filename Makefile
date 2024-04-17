.PHONY: build run all

all: build run

build:
	go build ./cli/main/main.go

run:
	./main

