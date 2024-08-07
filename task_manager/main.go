package main

import (
	"task_manager/router"
	"task_manager/data"
	"task_manager/storage"
	"task_manager/controllers"
	"github.com/gin-gonic/gin"
)

func main(){
	newstorage := Storage.NewNoSqlConnection()
	newservice := service.NewServices(newstorage)
	controller := Controllers.NewController(newservice)

	serverRoute := gin.Default()
	router.GetRoute(serverRoute, controller)
	serverRoute.Run()
}
