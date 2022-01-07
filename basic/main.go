package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type request struct {
	User string `json:"user"`
}

type response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func handleRequest(req request) (response, error) {
	return response{
		StatusCode: 200,
		Body:       fmt.Sprintf("hello %s", req.User),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
