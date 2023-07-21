package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates*/*")

	r.Static("/", "./public")
	r.POST("/clicked", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "clicked.tmpl", gin.H{
			"hello": "world",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
