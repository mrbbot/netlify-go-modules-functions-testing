package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
	"netlify-go-modules-functions-testing/lambdaify"
)

func main() {
	r := gin.Default()

	r.GET("/.netlify/functions/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	lambda.Start(lambdaify.Lambdaify(r))
}
