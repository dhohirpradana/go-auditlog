package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/olivere/elastic"
)

func GetESClient() (*elastic.Client, error) {
	esURL := goDotEnvVariable("ES_URL")

	client, err := elastic.NewClient(elastic.SetURL(esURL),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

func storeToES(request interface{}, response interface{}, data interface{}) {
	ctx := context.Background()
	client, err := GetESClient()

	req := Request{
		Method: request.(*http.Request).Method,
		URL:    request.(*http.Request).URL.String(),
		Header: fmt.Sprintf("%+v", request.(*http.Request).Header),
		Body:   fmt.Sprintf("%+v", request.(*http.Request).Body),
	}

	res := Response{
		StatusCode: response.(*http.Response).StatusCode,
		Header:     fmt.Sprintf("%+v", response.(*http.Response).Header),
		Data:       data,
	}

	if err != nil {
		fmt.Println("Error getting ES client:", err)
	}

	now := time.Now()

	requestResponse := RequestResponse{
		Request:   req,
		Response:  res,
		Timestamp: now,
	}

	_, err = client.Index().
		Index("auditlog").
		Type("sapujagad").
		BodyJson(requestResponse).
		Do(ctx)

	if err != nil {
		fmt.Println("Error storing request and response:", err)
	}

	fmt.Println("Request and response stored in ES")
}
