package handler

import "github.com/gin-gonic/gin"

func ListTarefas(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "GET tarefas",
	})
}