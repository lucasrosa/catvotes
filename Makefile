.PHONY: build clean deploy

build-lambda:
	env GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/votes cmd/votes/lambda/*

build-cli:
	env GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/votes cmd/votes/cli/*

test: 
	env GO111MODULE=on go test ./... -cover
	
clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
