package main

import (
    "context"
    "encoding/json"
    "net/http"
    "strings"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/sirupsen/logrus"
    "github.com/yourusername/project/internal/model"
    "github.com/yourusername/project/internal/parser"
    "github.com/yourusername/project/internal/log"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    logrus.Info("Received request")
    var req model.Request
    if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
        logrus.Error("Error parsing request body: ", err)
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusBadRequest,
            Body:       `{"error": "Invalid request body"}`,
        }, nil
    }

    htmlContent := parser.TextToHTML(req.PlainText) + req.HTML
    doc, err := parser.ParseAndFixHTML(htmlContent)
    if err != nil {
        logrus.Error("Error parsing HTML: ", err)
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

func main() {
    log.InitLogger()
    lambda.Start(handler)
}
