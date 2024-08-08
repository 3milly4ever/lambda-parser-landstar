package parser

import (
    "strings"
    "github.com/PuerkitoBio/goquery"
    "github.com/3milly4ever/lambda-parser-landstar/internal/model"
)

func TextToHTML(plainText string) string {
    return "<html><body>" + strings.ReplaceAll(plainText, "\n", "<br>") + "</body></html>"
}

func ConvertToHTML(plainText string) (string, error) {
    lines := strings.Split(plainText, "\n")
    var htmlBuilder strings.Builder

    htmlBuilder.WriteString("<html><body>")
    for _, line := range lines {
        formattedLine := fmt.Sprintf("<p>%s</p>", strings.TrimSpace(line))
        htmlBuilder.WriteString(formattedLine)
    }
    htmlBuilder.WriteString("</body></html>")

    return htmlBuilder.String(), nil
}


func ParseAndFixHTML(htmlContent string) (*goquery.Document, error) {
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
    if err != nil {
        return nil, err
    }
    return doc, nil
}

func ParseEmail(req model.Request) (*model.Response, error) {
    htmlContent := TextToHTML(req.PlainText) + req.HTML
    doc, err := ParseAndFixHTML(htmlContent)
    if err != nil {
        return nil, err
    }

    return &model.Response{
        StatusCode: 200,
        Body:       doc.Text(),
    }, nil
}
