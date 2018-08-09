package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHttp(config Config) {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.POST("/paste", func(c *gin.Context) {
		link := c.Query("link")
		c.JSON(http.StatusOK, gin.H{"shortLink": link})
	})

	r.GET("/:hash", func(c *gin.Context) {
		hash := c.Param("hash")

		c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("https://google.com?q=%s", hash))
	})

	r.Run(config.HttpAddr)
}
