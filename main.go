package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Channel struct {
	Name string
}

func main() {

	channels := []Channel{
		{
			Name: "test",
		},
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates*/*")

	r.Static("/public", "./public")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"channels": channels,
		})

	})
	r.POST("/channels", func(ctx *gin.Context) {
		channel := &Channel{}
		if err := ctx.Bind(channel); err != nil {
			return
		}
		channels = append(channels, *channel)
		ctx.HTML(http.StatusOK, "channels", gin.H{
			"channels": channels,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
