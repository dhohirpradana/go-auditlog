package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/olivere/elastic"
	"log"
	"net/http"
	"os"
	"testing"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func TestGetESClient(t *testing.T) {
	esURL := os.Getenv("ES_URL")

	fmt.Println(esURL)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}
	client, err := elastic.NewClient(
		elastic.SetHttpClient(c),
		elastic.SetURL(esURL),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetBasicAuth("elastic", "at2M0OSje04IU0Ei96lG931h"),
	)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	_, _, err = client.Ping(esURL).Do(context.Background())
	if err != nil {
		t.Fatalf("Failed to ping Elasticsearch: %v", err)
	}

	exists, err := client.IndexExists("auditlog").Do(context.Background())
	if err != nil {
		t.Fatalf("Failed to get Index: %v", err)
	}
	if !exists {
		t.Fatalf("Index does not exist yet.")
	}
}
