package utils

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/davejohnston/go-simple-websockets-chat-app/pkg/model"
)

var (
	tableName = os.Getenv("TABLE_NAME")
)

func getDB() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	return dynamodb.New(sess)
}

func getAPIGateway(request events.APIGatewayWebsocketProxyRequest) *apigatewaymanagementapi.ApiGatewayManagementApi {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return apigatewaymanagementapi.New(sess, aws.NewConfig().WithEndpoint(request.RequestContext.DomainName+"/"+request.RequestContext.Stage))
}

// StoreItem saves a connectionId to DynamoDB
func StoreItem(id string) error {

	item := model.Item{
		ConnectionID: id,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = getDB().PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

// DeleteItem deletes a connectionId item from DynamoDB
func DeleteItem(id string) error {

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"connectionId": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := getDB().DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}

// GetItem gets a connectionId item from DynamoDB
func GetItem(id string) (*model.Item, error) {

	result, err := getDB().GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"connectionId": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	item := model.Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// PostConnection sends the request payload back to the client using the ConnectionId as the identifier.
func PostConnection(item *model.Item, request events.APIGatewayWebsocketProxyRequest) error {

	gw := getAPIGateway(request)
	_, err := gw.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: &item.ConnectionID,
		Data:         []byte(request.Body),
	})
	if err != nil {
		return err
	}

	return nil
}
