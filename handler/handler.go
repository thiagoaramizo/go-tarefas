package handler

import "github.com/gin-gonic/gin"

func CreateTarefa(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"data": "POST tarefa",
	})
}

func ListTarefas(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"data": "GET tarefas",
	})
}

func GetTarefa(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"data": "GET tarefa",
	})
}

func UpdateTarefa(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"data": "PUT tarefa",
	})
}

func DeleteTarefa(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"data": "DELETE tarefa",
	})
}