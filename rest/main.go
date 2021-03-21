package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		fmt.Printf("menerima request dari %s\n", c.ClientIP())
		a, err := strconv.Atoi(c.PostForm("a"))
		if err != nil {
			fmt.Println("parameter invalid!")
			c.JSON(400, gin.H{
				"message": "invalid parameter",
			})
			return
		}
		b, err := strconv.Atoi(c.PostForm("b"))
		if err != nil {
			fmt.Println("parameter invalid!")
			c.JSON(400, gin.H{
				"message": "invalid parameter",
			})
			return
		}
		hasil := a + b
		fmt.Printf("hasil dari %d + %d adalah %d\n", a, b, hasil)
		c.JSON(200, gin.H{
			"a":     a,
			"b":     b,
			"hasil": hasil,
		})
	})
	r.Run()
}
