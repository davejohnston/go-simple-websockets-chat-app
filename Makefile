.PHONY: deps clean build

all: build

deps:
	go get -u ./...

clean:
	rm -rf ./cmd/connect/connect
	rm -rf ./cmd/disconnect/disconnect
	rm -rf ./cmd/sendMessage/sendMessage

build:
	GOOS=linux GOARCH=amd64 go build -o cmd/connect/connect ./cmd/connect/connect.go
	GOOS=linux GOARCH=amd64 go build -o cmd/disconnect/disconnect ./cmd/disconnect/disconnect.go
	GOOS=linux GOARCH=amd64 go build -o cmd/sendMessage/sendMessage ./cmd/sendMessage/sendMessage.go

package:
	sam package --template-file template.yaml --output-template-file packaged.yaml --s3-bucket daves-chatapp

deploy:
	sam deploy  --template-file packaged.yaml --stack-name go-simple-websockets-chat-app --capabilities CAPABILITY_IAM

publish: build package deploy