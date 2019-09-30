package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davejohnston/go-simple-websockets-chat-app/pkg/utils"
)

// Handler is the Lambda handler.
func Handler(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Received Disconnect request %s for client %s", request.RequestContext.RequestID, request.RequestContext.ConnectionID)

	err := utils.DeleteItem(request.RequestContext.ConnectionID)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	log.Println("Successfully deleted '" + request.RequestContext.ConnectionID)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)
}
