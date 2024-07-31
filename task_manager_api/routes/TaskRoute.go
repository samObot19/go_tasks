package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cntr "task_manager_api/controllers"
)

func StartRoute(){
	fmt.Println("Task manager API")
	router := gin.Default()
	router.GET("/tasks", cntr.GetTasks)
	router.GET("/tasks/:id", cntr.GetTask)
	router.DELETE("/tasks/:id", cntr.RemoveTask)
	router.PUT("/tasks/:id", cntr.UpdateTask)
	router.POST("/tasks", cntr.AddTask)
	router.Run()
}