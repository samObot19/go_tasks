package Controllers


import (
	"task_manager/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"net/http"
)


var jwtSecret = []byte("your_jwt_secret")


func(cnt  *Controller) Register(cnx *gin.Context){
	var user models.User
	err := cnx.ShouldBindJSON(&user)
	
	if err != nil{
		cnx.JSON(400, gin.H{"message" : "Invalid request payload!"})
		return 
	}
	
	_, ok := cnt.Service.GetUser(&user.UserName)
	
	if ok == nil{
		cnx.JSON(200,gin.H{"message" : "The User already registered!"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil{
		fmt.Println(err)
		cnx.JSON(500, gin.H{"message" : "Internal server error"})
		return
	}

	user.Password = string(hashedPassword)
	err = cnt.Service.AddUser(&user)
	

	if err != nil{
		fmt.Println(err)
		cnx.JSON(500, gin.H{"message" : "Internal server error!"})
		return
	}

	cnx.JSON(200, gin.H{"message" : "The user registered successfully!"})
}


func (cnt *Controller) Login(cnx *gin.Context){
	var user models.User

	if err := cnx.ShouldBindJSON(&user); err != nil{
		cnx.JSON(400, gin.H{"error" : "Invalid"})
	}

	loggedUser, err := cnt.Service.GetUser(&user.UserName)


	if err != nil{
		cnx.JSON(404, gin.H{"error" : err.Error()})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(loggedUser.Password), []byte(user.Password)) != nil {
		cnx.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  loggedUser.UserName,
		"role": loggedUser.Role,
		"ExpiresAt" : time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtToken, err := token.SignedString(jwtSecret)
	
	if err != nil {
		cnx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	cnx.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})

}

func(cnt *Controller) Promote(cnx *gin.Context){
	username := cnx.Param("username")
	err := cnt.Service.PromoteUser(username)

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

