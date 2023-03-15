package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func IndexHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{"message": "Hello!" + name})
	c.XML(200,Person{FirstName: "shadman", LastName: "sakib"})
}

func main() {
	router := gin.Default()
	//router.GET("/", IndexHandler)
	router.GET("/:name", IndexHandler)
	router.Run()

}
