package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST Opening",
	})
}
