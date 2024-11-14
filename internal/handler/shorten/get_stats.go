package shorten

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-shorten/internal/constant/errormsg"
	"go-shorten/internal/util"
)

func (h Handler) GetShortenStats(c *gin.Context) {
	shortCode := c.Param("code")
	cleanCode := util.RemoveNonAlphabets(shortCode)

	res, err := h.shortenService.GetStats(c, cleanCode)
	if err != nil {
		if errors.Is(err, errormsg.ErrShortenNotFound) {
			util.ErrResourceNotFoundResponse(c, err.Error())
			return
		}

		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success get shorten stats", res)
	return
}
