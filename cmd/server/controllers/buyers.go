package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBuyers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"value": "ok",
	})
}
