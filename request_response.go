package main

import "time"

type Request struct {
	Method string      `json:"method"`
	URL    string      `json:"url"`
	Header interface{} `json:"header"`
	Body   interface{} `json:"body"`
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Header     string      `json:"header"`
	Data       interface{} `json:"data"`
}

type RequestResponse struct {
	Request   interface{} `json:"request"`
	Response  interface{} `json:"response"`
	Timestamp time.Time   `json:"timestamp"`
}
