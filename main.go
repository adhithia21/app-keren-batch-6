package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200,
			gin.H{
				"status":      200,
				"version":     "2.0",
				"description": "Penambahan fitur abcd",
			})
	})
	r.Run(":8000")
}
