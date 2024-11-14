package shorten

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-shorten/internal/constant/errormsg"
	"go-shorten/internal/util"
	"net/http"
)

func (h Handler) GetShorten(c *gin.Context) {
	shortCode := c.Param("code")
	cleanCode := util.RemoveNonAlphabets(shortCode)

	shorten, err := h.shortenService.GetByCode(c, cleanCode)
	if err != nil {
		if errors.Is(err, errormsg.ErrShortenNotFound) {
			util.ErrResourceNotFoundResponse(c, err.Error())
			return
		}

		util.ErrInternalResponse(c, err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, shorten.OriginalURL)
	return
}
