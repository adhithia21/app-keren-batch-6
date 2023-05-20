package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200,
			gin.H{
				"status":  200,
				"version": "1.0",
			})
	})
	r.Run(":8000")
}
