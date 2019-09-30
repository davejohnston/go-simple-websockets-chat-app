.PHONY: deps clean build

NAME=go-simple-websockets-chat-app
BUCKET=$(NAME)

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
	@sam package --template-file template.yaml --output-template-file packaged.yaml --s3-bucket $(BUCKET)

deploy:
	@sam deploy  --template-file packaged.yaml --stack-name $(NAME) --capabilities CAPABILITY_IAM

info:
	@aws cloudformation describe-stacks --stack-name $(NAME) --query "Stacks[].Outputs[?OutputKey=='WebSocketURI'].OutputValue"

publish: build package deploy info

delete:
	aws cloudformation delete-stack --stack-name $(NAME)