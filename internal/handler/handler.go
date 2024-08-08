package main

import (
    "context"
    "encoding/json"
    "net/http"
    "github.com/aws/aws-lambda-go/events"
    "github.com/sirupsen/logrus"
    "github.com/3milly4ever/lambda-parser-landstar/internal/log"
    "github.com/3milly4ever/lambda-parser-landstar/internal/model"
    "github.com/3milly4ever/lambda-parser-landstar/internal/parser"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    logrus.Info("Received request")
    logrus.Infof("Request body: %s", request.Body)

    var req model.Request
    err := json.Unmarshal([]byte(request.Body), &req)
    if err != nil {
        logrus.Errorf("Error parsing request body: %v", err)
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusBadRequest,
            Body:       `{"error": "Invalid request body"}`,
        }, nil
    }

    logrus.Infof("Parsed request: %+v", req)

    htmlContent := parser.TextToHTML(req.PlainText) + req.HTML
    logrus.Infof("HTML Content: %s", htmlContent)

    doc, err := parser.ParseAndFixHTML(htmlContent)
    if err != nil {
        logrus.Errorf("Error parsing HTML: %v", err)
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       `{"error": "Error processing HTML"}`,
        }, nil
    }

    logrus.Info("Successfully parsed HTML")

    return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       doc.Text(),
    }, nil
}
