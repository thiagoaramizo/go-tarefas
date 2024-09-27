package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagoaramizo/go-tarefas/handler"
)

func initializeRoutes(router *gin.Engine){
	
	v1 := router.Group("/api/v1")
	{
		v1.GET("/tarefas", handler.ListTarefas )
		v1.POST("/tarefa", handler.CreateTarefa)
		v1.GET("/tarefa/:id", handler.GetTarefa)
		v1.PUT("/tarefas/:id", handler.UpdateTarefa)
		v1.DELETE("/tarefas/:id", handler.DeleteTarefa)
	}

}