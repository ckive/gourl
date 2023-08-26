package handler

import (
	"net/http"

	"github.com/ckive/gourl/backend/constants"
	"github.com/ckive/gourl/backend/shortener"
	"github.com/ckive/gourl/backend/store"

	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl    string `json:"long_url" binding:"required"`
	CustomLink string `json:"custom_link" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if custom link is in cache
	var shortUrl string
	if customLinkInCache := store.CustomLinkInCache(creationRequest.CustomLink); customLinkInCache {
		// In Cache,
		// TODO use frontend to do checks or something for now just give rando
		shortUrl = shortener.GenerateShortLink(creationRequest.LongUrl)
	} else {
		// not in cache, create the short url
		shortUrl = creationRequest.CustomLink[:min(constants.ShortLinkLength, len(creationRequest.CustomLink))]
	}
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.CustomLink)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
