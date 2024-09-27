package handler

import (
	"github.com/gin-gonic/gin"
)

func DeleteTarefa(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "DELETE tarefa",
	})
}