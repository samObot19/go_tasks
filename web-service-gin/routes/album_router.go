package routes

import (
	"web-service-gin/controllers"
    "github.com/gin-gonic/gin"
)

func StartRoute(){
	router := gin.Default()

    router.GET("/albums", controllers.GetAlbums)
    router.GET("/albums/:id", controllers.GetAlbumByID)
    router.POST("/albums", controllers.PostAlbums)

    router.Run("localhost:8080")
}