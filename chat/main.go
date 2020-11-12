package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./views/*.html")

	r.GET("/", index)

	r.Run(":8080")
}
