package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LinkCtrl struct {
	linkService LinkService
}

func (lc *LinkCtrl) ShortenLink(c *gin.Context) {
	link := c.Query("link")
	if link == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "link must be valid url",
		})
		return
	}
	clientIp := c.Request.RemoteAddr
	shortLink := lc.linkService.ShortenLink(link, clientIp)
	c.JSON(http.StatusOK, gin.H{"shortLink": shortLink})
}

func (lc *LinkCtrl) Redirect(c *gin.Context) {
	hash := c.Param("hash")
	link := lc.linkService.FindLink(hash)
	c.Redirect(http.StatusPermanentRedirect, link)
}
