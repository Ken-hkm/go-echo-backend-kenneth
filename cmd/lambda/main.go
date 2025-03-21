package main

import (
	"context"
	"github.com/Ken-hkm/go-echo-backend-kenneth/internal/db"
	"log"

	routes "github.com/Ken-hkm/go-echo-backend-kenneth/internal/api"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda

func init() {

	//conn to Mongo
	db.ConnectMongoDB()
	e := echo.New()
	routes.RegisterRoutes(e)

	// Wrap Echo with AWS Lambda adapter
	echoLambda = echoadapter.New(e)
}

// Lambda handler function
func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	log.Println("Starting AWS Lambda with Echo...")
	lambda.Start(handler)
}
