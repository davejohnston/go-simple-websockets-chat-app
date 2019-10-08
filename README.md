# go-simple-websockets-chat-app

[![Go Report Card](https://goreportcard.com/badge/github.com/davejohnston/go-simple-websockets-chat-app)](https://goreportcard.com/badge/github.com/davejohnston/go-simple-websockets-chat-app)

This is a golang implementation for the AWS example simple-websocket-chat-app.  See https://github.com/aws-samples/simple-websockets-chat-app for the original implementation.  There are three functions contained within the directories and a SAM template that wires them up to a DynamoDB table and provides the minimal set of permissions needed to run the app:


```bash
.
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── cmd
│   ├── connect
│   │   └── connect.go          <-- Lambda function code for connect
│   ├── disconnect
│   │   └── disconnect.go       <-- Lambda function code for disconnect
│   └── sendMessage
│       └── sendMessage.go      <-- Lambda function code for sendMessage
└── template.yaml               <-- SAM template for Lambda Functions and DDB

```

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Setup process

### Installing dependencies

In this example we use the built-in `go get` and the only dependency we need is AWS Lambda Go SDK:

```shell
go get -u github.com/aws/aws-lambda-go/...
```

**NOTE:** As you change your application code as well as dependencies during development, you might want to research how to handle dependencies in Golang at scale.

### Building

Golang is a statically compiled language, meaning that in order to run it you have to build the executable target.

You can issue the following command in a shell to build it:

```shell
	GOOS=linux GOARCH=amd64 go build -o cmd/connect/connect ./cmd/connect/connect.go
	GOOS=linux GOARCH=amd64 go build -o cmd/disconnect/disconnect ./cmd/disconnect/disconnect.go
	GOOS=linux GOARCH=amd64 go build -o cmd/sendMessage/sendMessage ./cmd/sendMessage/sendMessage.go
```

**NOTE**: If you're not building the function on a Linux machine, you will need to specify the `GOOS` and `GOARCH` environment variables, this allows Golang to build your function for another system architecture and ensure compatibility.

### Local development

First and foremost, we need a `S3 bucket` where we can upload our Lambda functions packaged as ZIP before we deploy anything - If you don't have a S3 bucket to store code artifacts then this is a good time to create one:

```bash
aws s3 mb s3://BUCKET_NAME
```

Next, run the following command to package our Lambda function to S3:

```bash
sam package \
    --output-template-file packaged.yaml \
    --s3-bucket REPLACE_THIS_WITH_YOUR_S3_BUCKET_NAME
```

Next, the following command will create a Cloudformation Stack and deploy your SAM resources.

```bash
sam deploy \
    --template-file packaged.yaml \
    --stack-name go-simple-websockets-chat-app \
    --capabilities CAPABILITY_IAM
```

### Testing 
Install the wscat client.
```bash
npm install -g wscat
```

Now make a request to your depolyed API gateway

```bash
wscat -c wss://aabbccddee.execute-api.us-east-1.amazonaws.com/Prod/
```

This will present a terminal, where you can send messages that will be echo'd back.  e.g. 
```
connected (press CTRL+C to quit)
> {"message":"sendmessage", "data":"hello world"}
< {"message":"sendmessage", "data":"hello world"}
```


> **See [Serverless Application Model (SAM) HOWTO Guide](https://github.com/awslabs/serverless-application-model/blob/master/HOWTO.md) for more details in how to get started.**

After deployment is complete you can run the following command to retrieve the API Gateway Endpoint URL:

```bash
aws cloudformation describe-stacks \
    --stack-name go-simple-websockets-chat-app \
    --query 'Stacks[].Outputs'
``` 

# Appendix

### Deleting Stack

```bash
aws cloudformation delete-stack --stack-name go-simple-websockets-chat-app
```
### Golang installation

Please ensure Go 1.x (where 'x' is the latest version) is installed as per the instructions on the official golang website: https://golang.org/doc/install

A quickstart way would be to use Homebrew, chocolatey or your linux package manager.

#### Homebrew (Mac)

Issue the following command from the terminal:

```shell
brew install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
brew update
brew upgrade golang
```

#### Chocolatey (Windows)

Issue the following command from the powershell:

```shell
choco install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
choco upgrade golang
```
