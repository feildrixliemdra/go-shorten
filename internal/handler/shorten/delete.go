package shorten

import (
	"github.com/gin-gonic/gin"
	"go-shorten/internal/util"
)

func (h Handler) Delete(c *gin.Context) {
	shortCode := c.Param("code")
	cleanCode := util.RemoveNonAlphabets(shortCode)

	err := h.shortenService.Delete(c, cleanCode)
	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success delete shorten", nil)
	return
}
