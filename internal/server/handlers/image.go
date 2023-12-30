package handlers

import (
	"gamelib/internal/actions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) PostImage(ctx *gin.Context) {
	img, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "image error",
		})
		return
	}

	imgName := ctx.Param("name")
	if len(imgName) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "image name error",
		})
		return
	}

	if err := ctx.SaveUploadedFile(img, actions.PathGrids+imgName); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
		})
		return
	}

	ctx.String(http.StatusOK, "Files uploaded")
}
