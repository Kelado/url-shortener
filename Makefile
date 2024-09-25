SVC_NAME=app

run: build
	./bin/$(SVC_NAME)

test: 
	go clean -testcache
	go test ./controllers/ ./repositories/ ./handlers/ -cover 

build:
	go build -o ./bin/$(SVC_NAME)

clean:
	rm ./bin/$(SVC_NAME)