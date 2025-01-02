package routes

import (
	"crud_project/dto"
	"crud_project/logger"
	"crud_project/models"
	"crud_project/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context){
      users, err := models.GetAllUsers()

	  if err!=nil{
		utility.SendInternalServerError(context)
		return
	  }
	  utility.SendSuccessResponse(context, 200, users)
}

func CreateUsers(context *gin.Context){
	var userObj models.User
	err := context.ShouldBindJSON(&userObj)
	logger.InfoLogger.Println("Working with create user function")

	if err!= nil{
		logger.ErrorLogger.Print("Error while reading the input")
		utility.SendBadrequest(context) 
		return 
	}
	Valerror := dto.CheckUserInput(userObj)
	if Valerror != nil{
		utility.SendValidationError(context, Valerror)
		return
	}
	err = userObj.Save()
	if err!= nil{
		logger.ErrorLogger.Print("Error in User creation")
		utility.SendInternalServerError(context)
		return
	}

	logger.InfoLogger.Println("User created successfully")
	utility.SendSuccessResponse(context, 201, userObj)
	
}

func Login(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err!= nil{
		logger.ErrorLogger.Print("Error while reading the input")
		utility.SendBadrequest(context) 
		return 
	}
	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "login sucessful!"})
}