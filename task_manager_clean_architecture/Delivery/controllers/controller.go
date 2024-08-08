package controller

import (
	"github.com/task_manager/usecases"
	"github.com/task_manager/Infrastructure"
	"github.com/task_manager/Domain"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

type Handler struct{
	task	*usecase.TaskUsecase
	user	*usecase.UserUsecase
}

func NewHandler(tk *usecase.TaskUsecase, ur *usecase.UserUsecase) *Handler{
	return &Handler{
			task : tk,
			user : ur,
		}
}

func (cnt *Handler)AddTask(ctx *gin.Context){
	var newTask domain.Task
	err := ctx.ShouldBindJSON(&newTask)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        	return
    	}

    	err = cnt.task.AddTask(&newTask)
	if err != nil{
		ctx.JSON(http.StatusCreated, gin.H{"message": err})
	}
    	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})

}


func (cnt *Handler) GetTask(ctx *gin.Context){
	id := ctx.Param("id")
	val, err := cnt.task.GetTask(&id)

    	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"message" : "Task not found"})
    	}else{
        	ctx.JSON(http.StatusOK, val)
    	}

}

func (cnt *Handler) GetTasks(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"tasks": cnt.task.GetTasks()})
}

func (cnt *Handler) RemoveTask(ctx *gin.Context){
	id := ctx.Param("id")

	err := cnt.task.RemoveTask(&id)

	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"message" : "Task not found"})
	}else{
		ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
	}

}

func (cnt *Handler) UpdateTask(ctx *gin.Context){
	id := ctx.Param("id")
	var updatedTask domain.Task

	err := ctx.ShouldBindJSON(&updatedTask)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cnt.task.UpdateTask(&id, &updatedTask)

	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}else{
		ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
	}
}


var jwtSecret = "your_jwt_secret"


func(cnt  *Handler) Register(cnx *gin.Context){
	var user domain.User
	err := cnx.ShouldBindJSON(&user)
	
	if err != nil{
		cnx.JSON(400, gin.H{"message" : "Invalid request payload!"})
		return 
	}
	
	_, ok := cnt.user.GetUser(&user.UserName)
	
	if ok == nil{
		cnx.JSON(200,gin.H{"message" : "The User already registered!"})
		return
	}

	hashedPassword, err := Infrastructure.EncryptPassword(jwtSecret)
	if err != nil{
		fmt.Println(err)
		cnx.JSON(500, gin.H{"message" : "Internal server error"})
		return
	}

	user.Password = hashedPassword
	err = cnt.user.AddUser(&user)
	

	if err != nil{
		fmt.Println(err)
		cnx.JSON(500, gin.H{"message" : "Internal server error!"})
		return
	}

	cnx.JSON(200, gin.H{"message" : "The user registered successfully!"})
}


func (cnt *Handler) Login(cnx *gin.Context){
	var user domain.User
	jwt := Infrastructure.NewJWTService(jwtSecret)

	if err := cnx.ShouldBindJSON(&user); err != nil{
		cnx.JSON(400, gin.H{"error" : "Invalid"})
	}

	loggedUser, err := cnt.user.GetUser(&user.UserName)


	if err != nil{
		cnx.JSON(404, gin.H{"error" : err.Error()})
		return
	}

	if Infrastructure.IsValidPassword(user.Password, loggedUser.Password){
		cnx.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}

	jwtToken, err := jwt.GenerateToken(&loggedUser)
	
	if err != nil {
		cnx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	cnx.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})

}

func(cnt *Handler) Promote(cnx *gin.Context){
	username := cnx.Param("username")
	err := cnt.user.PromoteUser(username)

	if err != nil{
		cnx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	cnx.JSON(http.StatusOK,  gin.H{
		"message" : fmt.Sprintf("the user %s is promoted as an admin sucessfully!", username),
	})
}

