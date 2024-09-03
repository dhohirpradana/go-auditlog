package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetESClient() (*elastic.Client, error) {
	esURL := os.Getenv("ES_URL")
	esUser := os.Getenv("ES_USER")
	esPass := os.Getenv("ES_PASS")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	client, err := elastic.NewClient(
		elastic.SetHttpClient(c),
		elastic.SetURL(esURL),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetBasicAuth(esUser, esPass),
	)

	fmt.Println("ES initialized...")

	return client, err
}

func storeToES(request interface{}, response interface{}, data interface{}) {
	ctx := context.Background()
	client, err := GetESClient()
	esIndex := os.Getenv("ES_INDEX")

	req := Request{
		Method: request.(*http.Request).Method,
		URL:    request.(*http.Request).URL.String(),
		Header: request.(*http.Request).Header,
		Body:   request.(*http.Request).Body,
	}

	headersJSON, _ := json.Marshal(request.(*http.Request).Header)

	req.Header = string(headersJSON)

	bodyJSON, _ := json.Marshal(request.(*http.Request).Body)

	req.Body = string(bodyJSON)

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

	log.Printf("RR: %+v\n", requestResponse)
	//log.Printf("Request: %+v\n", requestResponse.Request)
	//log.Printf("Response: %+v\n", requestResponse.Response)

	_, err = client.Index().
		Index(esIndex).
		BodyJson(requestResponse).
		Do(ctx)

	if err != nil {
		fmt.Println("Error storing request and response:", err.Error())
	} else {
		fmt.Println("Request and response stored in ES")
	}
}
