package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const tableName = "Routes"

type Route struct {
	ID			int32
	Name   		string
	Length 		float64
	Description string
	Coordinates string
}

func routeHandler(w http.ResponseWriter, r *http.Request, routeRequest string) {

	if routeRequest == "all" {
		routes := getRoutes()
		jsonData, _ := json.Marshal(routes)
		fmt.Fprintf(w, "%s", jsonData)
	} else {
		route := getRouteById(routeRequest)
		jsonData, _ := json.Marshal(route)
		fmt.Fprintf(w, "%s", jsonData)
	}
	
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

func getRoutes() []Route {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	
	svc := dynamodb.New(sess)

	proj := expression.NamesList(expression.Name("ID"), expression.Name("Name"), expression.Name("Length"))

	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	routes := make([]Route, len(result.Items))
	for _, i := range result.Items {
		route := Route{}

		err = dynamodbattribute.UnmarshalMap(i, &route)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		routes = append(routes, route)
	}

	return routes
}
