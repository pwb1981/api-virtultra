package route

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
)

const tableName = "Routes"

type Route struct {
	ID			int32
	Name   		string
	Length 		float64
	Description string
	Coordinates string
}

func getRouteById(routeId string) Route {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	
	svc := dynamodb.New(sess)

	route := Route{}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(routeId),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return route
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &route)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if route.Name == "" {
		fmt.Println("Could not find '" + routeId + ")")
		return route
	}

	return route
}
