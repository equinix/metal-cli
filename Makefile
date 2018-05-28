default: build generete-docs

build:
	GOOS=linux go build -o bin/linux/packet
	GOOS=darwin go build -o bin/darwin/packet

clean: 
	rm -rf bin/

generete-docs: 
	GENDOCS=true go run main.go