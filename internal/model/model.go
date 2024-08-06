package model

type Request struct {
	PlainText string `json:"plain_text"`
	HTML      string `json:"html"`
	Data      string `json:"data"`
}

type Response struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"body"`
}

