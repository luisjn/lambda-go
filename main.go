package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-redis/redis/v8"
)

var (
	ctx        = context.Background()
	redis_host = os.Getenv("REDIS_HOST")
	redis_port = 6379
	rdb        = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redis_host, redis_port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
)

type request struct {
	Name string `json:"name"`
}

type response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func handleRequest(req request) (response, error) {
	fmt.Printf("Recieved request: %v", req)
	if req.Name != "" {
		err := rdb.Set(ctx, "name", req.Name, 0).Err()
		if err != nil {
			return response{
				StatusCode: 500,
				Body:       fmt.Sprintf("Oops! Could not write %s to Redis", req.Name),
			}, err
		}
		return response{
			StatusCode: 200,
			Body:       fmt.Sprintf("Success! %s was written to Redis", req.Name),
		}, nil
	}

	name, err := rdb.Get(ctx, "name").Result()
	if err == redis.Nil {
		return response{
			StatusCode: 200,
			Body:       "Oops! key does not exist",
		}, err
	} else if err != nil {
		return response{
			StatusCode: 500,
			Body:       "Oops! Could not read the key name in Redis",
		}, err
	}

	return response{
		StatusCode: 200,
		Body:       fmt.Sprintf("Hello %s nice to meet you", name),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
