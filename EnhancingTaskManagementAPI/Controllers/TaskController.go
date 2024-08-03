package Controllers

import (
	"EnhancingTaskManagementAPI/Services"
	"EnhancingTaskManagementAPI/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
    Service *Service.Services
}

func NewController(newService *Service.Services) *Controller {
    return &Controller{Service: newService}
}

func (cnt *Controller)AddTask(ctx *gin.Context){
	var newTask Models.Task
	err := ctx.ShouldBindJSON(&newTask)

    if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    cnt.Service.AddTask(&newTask)
    ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})

}


func (cnt *Controller) GetTask(ctx *gin.Context){
	id := ctx.Param("id")
	val, err := cnt.Service.GetTask(&id)

    if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"message" : "Task not found"})
    }else{
        ctx.JSON(http.StatusOK, val)
    }

}

func (cnt *Controller) GetTasks(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"tasks": cnt.Service.GetTasks()})
}

func (cnt *Controller) RemoveTask(ctx *gin.Context){
	id := ctx.Param("id")

	err := cnt.Service.RemoveTask(&id)

	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"message" : "Task not found"})
	}else{
		ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
	}

}

func (cnt *Controller) UpdateTask(ctx *gin.Context){
	id := ctx.Param("id")
	var updatedTask Models.Task

	err := ctx.ShouldBindJSON(&updatedTask)

	if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	err = cnt.Service.UpdateTask(&id, &updatedTask)

	if err != nil{
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}else{
			ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
	}

}