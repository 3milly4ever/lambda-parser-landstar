package parser

import (
    "strings"
    "github.com/PuerkitoBio/goquery"
)

func TextToHTML(plainText string) string {
    return "<html><body>" + strings.ReplaceAll(plainText, "\n", "<br>") + "</body></html>"
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
