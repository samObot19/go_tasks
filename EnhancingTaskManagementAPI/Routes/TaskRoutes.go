package Routes

import (
        "fmt"
        "github.com/gin-gonic/gin"
        "EnhancingTaskManagementAPI/Controllers"
)

func GetRoute(router *gin.Engine, cntr *Controllers.Controller){
        fmt.Println("Task manager API")
        router.GET("/tasks", cntr.GetTasks)
        router.GET("/tasks/:id", cntr.GetTask)
        router.DELETE("/tasks/:id", cntr.RemoveTask)
        router.PUT("/tasks/:id", cntr.UpdateTask)
        router.POST("/tasks", cntr.AddTask)
}