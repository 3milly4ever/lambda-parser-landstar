package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/3milly4ever/lambda-parser-landstar/internal/log"
	"github.com/3milly4ever/lambda-parser-landstar/internal/handler"
)

func main(){
	log.InitLogger()
    lambda.Start(handler.HandleRequest)
}
