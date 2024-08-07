package router


import (
    "github.com/gin-gonic/gin"
    "task_manager/controllers"
	"task_manager/middleware"
)

func GetRoute(router *gin.Engine, cntr *Controllers.Controller){
	public := router.Group("/api")
    	{
        	public.POST("/login", cntr.Login)
        	public.POST("/register", cntr.Register)
    	}

    	protected := router.Group("/api")
    	{
        	protected.Use(middleware.AuthMiddleware())
        	protected.GET("/tasks", cntr.GetTasks)
        	protected.GET("/tasks/:id", cntr.GetTask)
        	admin := protected.Group("/admin")
        
        admin.Use(middleware.RoleBasedMiddleware("Admin"))
        {
            	admin.PUT("/tasks/:id", cntr.UpdateTask)
            	admin.DELETE("/tasks/:id", cntr.RemoveTask)
            	admin.POST("/tasks", cntr.AddTask)
            	admin.PUT("/users/promote/:username", cntr.Promote)
        }
    }
}

