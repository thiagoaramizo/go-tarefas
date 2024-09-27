package handler

import (
	"github.com/gin-gonic/gin"
)

func UpdateTarefa(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"data": "PUT tarefa",
	})
}