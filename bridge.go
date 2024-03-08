package main

import (
	"encoding/json"
	"fmt"
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
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	statusCode := response.StatusCode

	c.JSON(statusCode, data)

	storeToES(c.Request, response, data)
}
