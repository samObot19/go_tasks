
package router


import (
    "github.com/gin-gonic/gin"
    "github.com/task_manager/Delivery/controllers"
	"github.com/task_manager/Infrastructure"
)

func GetRoute(router *gin.Engine, cntr *controller.Handler){
	public := router.Group("/api")
    	{
        	public.POST("/login", cntr.Login)
        	public.POST("/register", cntr.Register)
    	}

    	protected := router.Group("/api")
    	{
        	protected.Use(Infrastructure.AuthMiddleware())
        	protected.GET("/tasks", cntr.GetTasks)
        	protected.GET("/tasks/:id", cntr.GetTask)
        	admin := protected.Group("/admin")

        	admin.Use(Infrastructure.RoleBasedMiddleware("Admin"))
        	{
            	admin.PUT("/tasks/:id", cntr.UpdateTask)
            	admin.DELETE("/tasks/:id", cntr.RemoveTask)
            	admin.POST("/tasks", cntr.AddTask)
            	admin.PUT("/users/promote/:username", cntr.Promote)
        	}
    }
}
