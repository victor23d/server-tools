package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var (
	value float32
)

func main() {

	port := 80
	value = 10
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"value": value,
		})
	})
	r.GET("/metrics", func(c *gin.Context) {
		s := fmt.Sprintf("# HELP http_requests_total The amount of requests served by the server in total\n"+
			"# TYPE http_requests_total counter\n"+
			"http_requests_total %f", value)
		c.String(http.StatusOK, "%s", s)
	})

	r.POST("/", handleV)
	r.Run(":" + strconv.Itoa(port))
}

func handleV(c *gin.Context) {
	type postValue struct {
		Value float32 `json:"value"`
	}
	var v postValue
	if err := c.ShouldBind(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value = v.Value
	log.Printf("value: %f\n", value)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
