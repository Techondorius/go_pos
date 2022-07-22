package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_pos/controller"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "good"})
	})

	price := r.Group("/api/price")
	{
		price.GET("/", controller.GetPrice)
		price.GET("/testdata", controller.InsertTestData)

	}

	r.Run()

	a := burger{100}
	fmt.Println(a)
}

type burger struct {
	price int
}