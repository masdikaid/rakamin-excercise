lint:
	go fmt ./...

build:
	go mod vendor
	go build -v -o bin/excercise-rakamin main.go