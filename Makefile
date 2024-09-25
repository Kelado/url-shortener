SVC_NAME=app

run:
	go run .

test: 
	go clean -testcache
	go test ./controllers/ ./repositories/ ./handlers/ -cover 

build:
	go build -o ./bin/$(SVC_NAME)

clean:
	rm ./bin/$(SVC_NAME)