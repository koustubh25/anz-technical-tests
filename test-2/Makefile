all: test build 

test: 
	CGO_ENABLED=0 go test

build: 
	CGO_ENABLED=0  GOOS=linux go build -o golang-version	