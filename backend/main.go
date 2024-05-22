package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// init items with those items: [a,b,c,d]
var items = []string{"a", "b", "c", "d"}

func main() {
	router := gin.Default()
	router.GET("/items", getItem)

	router.Run()
}

func getItem(c *gin.Context) {
	// return items as json
	c.IndentedJSON(http.StatusOK, items)
}
