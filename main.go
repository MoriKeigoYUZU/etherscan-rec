package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Transaction struct {
	Value string `json:"value"`
	From  string `json:"from"`
	To    string `json:"to"`
	Time  string `json:"timeStamp"`
}

type Result struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []Transaction `json:"result"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		contractAddress := "0xBB0E17EF65F82Ab018d8EDd776e8DD940327B28b"

		apiEndpoint := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=tokentx&address=%s&startblock=0&endblock=999999999&apikey=YourApiKeyToken", contractAddress)

		resp, err := http.Get(apiEndpoint)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var result Result
		err = json.Unmarshal(body, &result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response := result
		c.JSON(http.StatusOK, response)
	})

	router.Run(":8080")
}
