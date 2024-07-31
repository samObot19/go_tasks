package controllers

import (
	"task_manager_api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task Manager Project", Description: "Add/View/Delete Tasks", DueDate: time.Now(), Status: "In Progress"},
	{ID: "2", Title: "Books Management Project", Description: "Add/View/Delete Books", DueDate: time.Now().AddDate(0, 0, -1), Status: "Completed"},
}

func GetTasks(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTask(ctx *gin.Context){
	id := ctx.Param("id")

	for _, val := range tasks{
		if val.ID == id{
			ctx.JSON(http.StatusOK, val)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message" : "Task not found"})
}

func RemoveTask(ctx *gin.Context){
	id := ctx.Param("id")

	for i, val := range tasks{
		if val.ID == id{
			tasks = append(tasks[: i], tasks[i + 1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message" : "The task not found!"})
}

func UpdateTask(ctx *gin.Context){
	id := ctx.Param("id")
	var updatedTask models.Task

	err := ctx.ShouldBindJSON(&updatedTask)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks{
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func AddTask(ctx *gin.Context){
	var newTask models.Task

	err := ctx.ShouldBindJSON(&newTask)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}