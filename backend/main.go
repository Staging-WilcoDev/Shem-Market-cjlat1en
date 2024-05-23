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

// create a variable that store those itemm name:  'Galactic Goggles', 'Meteor Muffins', 'Alien Antenna Kit', 'Starlight Lantern', and 'Quantum Quill' and an id
var items = []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}{
	{1, "Galactic Goggles"},
	{2, "Meteor Muffins"},
	{3, "Alien Antenna Kit"},
	{4, "Starlight Lantern"},
	{5, "Quantum Quill"},
}

func getItem(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}
