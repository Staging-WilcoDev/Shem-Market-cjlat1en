package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.HEAD("/healthcheck", healthcheck)
	router.POST("/item", addItem)
	router.GET("/item/:id", getItemByID)

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
	Views int    `json:"views"`
}{
	{1, "Galactic Goggles", 0},
	{2, "Meteor Muffins", 0},
	{3, "Alien Antenna Kit", 0},
	{4, "Starlight Lantern", 0},
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, removeViews(items))
}

func removeViews(items []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Views int    `json:"views"`
}) []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
} {
	var itemsWithoutViews []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	for _, item := range items {
		itemsWithoutViews = append(itemsWithoutViews, struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{item.ID, item.Name})
	}
	return itemsWithoutViews
}

func addItem(c *gin.Context) {
    var item struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
		Views int    `json:"views"`
    }
    if err := c.BindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    item.ID = len(items) + 1
    item.Views = 0
    items = append(items, item)
    c.JSON(http.StatusCreated, gin.H{
		"id":   item.ID,
		"name": item.Name,
	})
}

func getItemByID(c *gin.Context) {
	id := c.Param("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var foundItem *struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Views int    `json:"views"`
	}
	for i := range items {
		if items[i].ID == itemID {
			foundItem = &items[i]
			break
		}
	}

	if foundItem == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	foundItem.Views++

	c.JSON(http.StatusOK, gin.H{
		"id":   foundItem.ID,
		"name": foundItem.Name,
	})
}