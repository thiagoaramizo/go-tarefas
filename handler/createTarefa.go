package handler

import "github.com/gin-gonic/gin"

func CreateTarefa(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"data": "POST tarefa",
	})
}