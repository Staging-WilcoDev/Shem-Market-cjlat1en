package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.HEAD("/healthcheck", healthcheck)
	router.POST("/item", addItem)

	router.Run()
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// create a variable that store those itemm name:  'Galactic Goggles', 'Meteor Muffins', 'Alien Antenna Kit', 'Starlight Lantern', and 'Quantum Quill' and an id
var items = []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}{
	{1, "Galactic Goggles"},
	{2, "Meteor Muffins"},
	{3, "Alien Antenna Kit"},
	{4, "Starlight Lantern"},
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func addItem(c *gin.Context) {
    var item struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    }
    if err := c.BindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    item.ID = len(items) + 1
    items = append(items, item)
    c.JSON(http.StatusCreated, item)
}