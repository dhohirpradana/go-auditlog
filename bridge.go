package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func bridgeRequest(c *gin.Context) {
	var err error
	var client = &http.Client{}
	var data interface{}

	protocol := c.Param("protocol")
	url := c.Param("url")
	path := c.Param("path")

	fullURL := protocol + url + "/" + path

	fmt.Println("Full URL:", fullURL)

	request, err := http.NewRequest(c.Request.Method, fullURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response, err := client.Do(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	statusCode := response.StatusCode

	dataInterface := convertToObject(data)

	fmt.Println(dataInterface)

	c.JSON(statusCode, dataInterface)

	storeToES(c.Request, response, dataInterface)
}

func convertToObject(data interface{}) map[string]interface{} {
	switch v := data.(type) {
	case string:
		return map[string]interface{}{
			"message": v,
		}
	case map[string]interface{}:
		return v
	default:
		return map[string]interface{}{
			"message": v,
		}
	}
}
