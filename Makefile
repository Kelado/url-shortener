SVC_NAME=app

run: build
	@./bin/$(SVC_NAME)

build:
	@go build -o ./bin/$(SVC_NAME)

test: 
	go clean -testcache
	go test ./controllers/ ./repositories/ ./handlers/ -cover 

clean:
	rm -f bin/*
	rm -f ./**/*.db