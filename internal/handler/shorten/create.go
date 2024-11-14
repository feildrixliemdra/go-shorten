package shorten

import (
	"github.com/gin-gonic/gin"
	"go-shorten/internal/payload"
	"go-shorten/internal/util"
)

func (h Handler) CreateShorten(c *gin.Context) {
	req := payload.CreateShortenRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	res, err := h.shortenService.Create(c, req)
	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success create shorten", res)
	return
}
