package handler

import "github.com/gin-gonic/gin"

func GetTarefa(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"data": "GET tarefa",
	})
}