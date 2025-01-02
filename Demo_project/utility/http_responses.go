// Contains helper function for various functions
package utility

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type str_map map[string]interface{} 

type OutputSrtuct struct{
	Status int        //`json: "status code"`
	Message string    //`json: "message"`
	Data interface{}  //`json:"data,omitempty"`

}

//Error messages corresponding to status code
const (
	InternalServerError = "Internal Server Error. Please try again later."
	BadRequest          = "Bad Request. Invalid input."
	Unauthorized        = "Unauthorized access."
	Forbidden           = "Forbidden access."
	NotFound            = "Resource not found."
	ServiceUnavailable  = "Service is unavailable. Please try again later."
)

// function to handle success 
func SendSuccessResponse(context *gin.Context, statusCode int, data ...interface{})  {
	promptMap := make(str_map)
	if len(data) >0 { 
		promptMap["data"] = data
	}else{
		promptMap["data"] = "No Content"
	}
	context.JSON(http.StatusOK, OutputSrtuct{
		Status: statusCode,
		Message: "Success",
		Data: promptMap ,
	} )
}

func SendInternalServerError(context *gin.Context){
	context.JSON(http.StatusInternalServerError, OutputSrtuct{
		Status: http.StatusInternalServerError,
		Message: InternalServerError,
	} )
}
func SendBadrequest(context *gin.Context){
	context.JSON(http.StatusBadRequest, OutputSrtuct{
		Status: http.StatusBadRequest,
		Message: BadRequest,
	} )
}
func SendValidationError(context *gin.Context, message []error){
	fmt.Println(message, len(message))
	promptMap := make(str_map)
	if len(message) > 0 {
		var errorMessages []string
		for _, err := range message {
			if err != nil {
				errorMessages = append(errorMessages, err.Error())
			}
		}
		promptMap["Error"] = errorMessages
	}
	context.JSON(http.StatusBadRequest, OutputSrtuct{
		Status: http.StatusBadRequest,
		Message: BadRequest,
		Data: promptMap, 
	} )
}

func SendNotFound(context *gin.Context){
	context.JSON(http.StatusNotFound, OutputSrtuct{
		Status: http.StatusNotFound,
		Message: NotFound,
	} )
}
func SendCustomError(context *gin.Context, statusCode int, errMessage string, data ...interface{}){
	promptError := make(str_map)

	if len(data) > 0 { 
		promptError["details"] = data
	}
	
	context.JSON(http.StatusUnauthorized, OutputSrtuct{
		Status: statusCode,
		Message: errMessage,
		Data: promptError ,
	} )

}



