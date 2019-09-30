package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davejohnston/go-simple-websockets-chat-app/pkg/utils"
)

// Handler is your Lambda function handler
func Handler(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Received SendMessage ID %s for client %s", request.RequestContext.RequestID, request.RequestContext.ConnectionID)

	item, err := utils.GetItem(request.RequestContext.ConnectionID)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	err = utils.PostConnection(item, request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "data sent",
	}, nil

}

func main() {
	lambda.Start(Handler)
}
