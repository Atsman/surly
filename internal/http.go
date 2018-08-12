package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHttp(config Config, linkCtrl LinkCtrl) {
	r := gin.Default()
	r.LoadHTMLGlob("../web/template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Surly url shorthand service",
		})
	})

	v1 := r.Group("/api/v1")
	v1.POST("/shorten", linkCtrl.ShortenLink)
	r.GET("/:hash", linkCtrl.Redirect)

	r.Run(config.HttpAddr)
}
