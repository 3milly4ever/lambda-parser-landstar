package parser

import (
	"fmt"
	"html"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ConvertToHTML(plainText string) (string, error) {
	lines := strings.Split(plainText, "\n")
	var htmlBuilder strings.Builder

	htmlBuilder.WriteString("<html><body>")
	for _, line := range lines {
		// Trim spaces and convert tabs to spaces
		formattedLine := strings.TrimSpace(strings.ReplaceAll(line, "\t", " &nbsp; &nbsp; "))

		// Only add a paragraph if the line is not empty
		if len(formattedLine) > 0 {
			htmlBuilder.WriteString(fmt.Sprintf("<p>%s</p>", html.EscapeString(formattedLine)))
		}
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
