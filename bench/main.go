package main

import (
	"net/http"
	"strconv"

	"bench/benchs"
	"github.com/gin-gonic/gin"
)

func main() {

	// alloc 100mb at start up
	benchs.MakeBlock(100)

	port := 80
	r := gin.Default()
	r.GET("/prime", func(c *gin.Context) {
		//c.JSON(200, gin.H{
		//	"value": pb,
		//)}
		c.String(http.StatusOK, "%+v\n", benchs.Pb)
	})
	r.POST("/prime", benchs.HandleP)

	r.GET("/memory", func(c *gin.Context) {
		//c.JSON(200, gin.H{
		//	"value": pb,
		//)}
		c.String(http.StatusOK, "%+v\n", benchs.GetMemUsage())
	})
	r.POST("/memory", benchs.HandleM)

	r.Run(":" + strconv.Itoa(port))

	defer benchs.CleanUpMem()
}
