package main

import (
	"EnhancingTaskManagementAPI/Controllers"
	"EnhancingTaskManagementAPI/Routes"
	"EnhancingTaskManagementAPI/Storage"
	"EnhancingTaskManagementAPI/Services"
	"github.com/gin-gonic/gin"
)

func main(){
	storage := Storage.NewNoSqlConnection()
	service := Service.NewServices(storage)
	controller := Controllers.NewController(service)

	serverRoute := gin.Default()
	Routes.GetRoute(serverRoute, controller)
	serverRoute.Run()
}
