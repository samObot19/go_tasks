package main

import (
	"github.com/gin-gonic/gin"
	"github.com/task_manager/usecase"
	"github.com/task_manager/Delivery/routers"
	"github.com/task_manager/Delivery/controllers"
	"github.com/task_manager/Repositories/database/mongodb"
)


func main(){
	task := mongodb.NewMongoTaskRepo()
	user := mongodb.NewMongoUserRepo()
	
	task_usecase := usecase.NewTaskUsecase(task)
	user_usecase := usecase.NewUserUsecase(user)

	cont := controller.NewHandler(task_usecase, user_usecase)
	serverRoute := gin.Default()
	router.GetRoute(serverRoute, cont)
	serverRoute.Run()
}


