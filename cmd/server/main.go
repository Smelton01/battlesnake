package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-chi/chi/v5"
	"github.com/smelton01/battlesnake/internal/api"
)

func main() {
	Execute()
}

func Execute() {
	log.Println("Started::: ")

	r := chi.NewRouter()
	api.BindAll(r)

	lambda.Start(httpadapter.New(r).ProxyWithContext)

}
