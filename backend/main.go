package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/items", getItem)

	router.Run()
}

// store items a,b,c and d in an array
var items = []string{"a", "b", "c", "d"}

func getItem(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}
