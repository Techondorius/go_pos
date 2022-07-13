package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "good"})
	})

	r.Run()

	a := burger{100}
	fmt.Println(a)
}

type burger struct {
	price int
}