package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

// type Person struct {
// 	XMLName   xml.Name `xml:"person"`
// 	FirstName string   `xml:"firstName,attr"`
// 	LastName  string   `xml:"lastName,attr"`
// }

// func IndexHandler(c *gin.Context) {
// 	name := c.Param("name")
// 	c.JSON(200, gin.H{"message": "Hello!" + name})
// 	c.XML(200,Person{FirstName: "shadman", LastName: "sakib"})
// }

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)

}

func main() {
	router := gin.Default()
	//router.GET("/", IndexHandler)
	//router.GET("/:name", IndexHandler)
	router.POST("/recipes", NewRecipeHandler)
	router.Run()

}
