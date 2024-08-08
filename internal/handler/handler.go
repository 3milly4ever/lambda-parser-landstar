package handler

import (
    "context"
    "encoding/json"
    "net/http"
    "github.com/aws/aws-lambda-go/events"
    "github.com/sirupsen/logrus"
    "github.com/3milly4ever/lambda-parser-landstar/internal/model"
    "github.com/3milly4ever/lambda-parser-landstar/internal/parser"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    logrus.Info("Received request")
    logrus.Info("Request body: ", request.Body)

    var req model.Request
    if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
        logrus.Error("Error parsing request body: ", err)
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusBadRequest,
            Body:       `{"error": "Invalid request body"}`,
        }, nil
    }

    htmlContent, err := parser.ConvertToHTML(req.PlainText)
    if err != nil {
        logrus.Error("Error converting text to HTML: ", err)
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       `{"error": "Error processing text"}`,
        }, nil
    }

    logrus.Info("Successfully converted text to HTML")

    return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       htmlContent,
    }, nil
}