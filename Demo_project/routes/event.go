package routes

import (
	"crud_project/dto"
	"crud_project/logger"
	"crud_project/models"
	"crud_project/utility"
	"strconv"
	"github.com/gin-gonic/gin"
)


func GetEvents(context *gin.Context){
	// function to get all the events from the database.
	event, err := models.GetAllEvents()
	if err != nil{
		utility.SendInternalServerError(context)
		return
	}
	utility.SendSuccessResponse(context, 200, event)
}

func CreateEvent(context *gin.Context) {
	// Function to create event, post request for event creation.

	var event models.Event  // Variable of type models.Event

	err := context.ShouldBindJSON(&event) // to read data from the request, and map it to the event struct

	logger.InfoLogger.Println("Working with create event function")
	if err!= nil{
		logger.ErrorLogger.Print("Error while reading the input")
		utility.SendBadrequest(context) 
		return 
	}
	logger.InfoLogger.Println("Running validation on the input data")

	Valerror := dto.CheckInput(event)
	if Valerror != nil {
		utility.SendValidationError(context, Valerror)

		return
	}
	err = event.Save() 

	if err != nil{
		logger.ErrorLogger.Print("Error in event creation")
		utility.SendInternalServerError(context)
		return
	}

	logger.InfoLogger.Println("Event created successfully")
	utility.SendSuccessResponse(context, 201, event)
}

func GetEvent(context *gin.Context) { 
	// gets single event by id

	eventid, err := strconv.ParseInt(context.Param("id"),10,64) // fetch id from path
	if err != nil{
		logger.ErrorLogger.Print("Unable to parse request")
		utility.SendBadrequest(context)
		return
	}

	event, err := models.GetEventByID(eventid)
	if err != nil{
		logger.ErrorLogger.Print("Unable to fetch event")
		utility.SendInternalServerError(context)
		return
	} 

		utility.SendSuccessResponse(context, 200, event)
}

func UpdateEvent(context *gin.Context) {
	// Function to update event details
	eventid, err := strconv.ParseInt(context.Param("id"),10,64) // fetch id from path
	if err != nil{
		logger.ErrorLogger.Print("Unable to parse request")
		// context.JSON(http.StatusBadRequest, utility.GetPrompt("Could not parse request data."))
		
		utility.SendBadrequest(context)
		return
	}
	_, err = models.GetEventByID(eventid) // check if event with the given id is there or not
	// data := utility.GetPrompt("event does not exist")
	if err != nil{
		// context.JSON(http.StatusInternalServerError, data)
		utility.SendInternalServerError(context)
		return
	}

	var updatedEvent models.Event    // new variable to store data send in the request by user
	err = context.ShouldBindJSON(&updatedEvent)   // extracting data from request and storing in updatedEvent
	
	if err!=nil{
		utility.SendBadrequest(context)
		return
	}
	updatedEvent.ID = eventid // copying id in the current event, so that new id is not created
	err = updatedEvent.Update()  // calling the update method to push data to db.

	if err !=nil {
		utility.SendInternalServerError(context)
		return
	}
	// context.JSON(http.StatusOK,utility.GetPrompt("event updated sucessfully"))
	utility.SendSuccessResponse(context, 200)
}

func DeleteEvent(context *gin.Context)  {
	// Function to delete event.
	eventId, err := strconv.ParseInt(context.Param("id"),10, 64)

	if err != nil{
		logger.ErrorLogger.Print("Unable to parse request")
		utility.SendBadrequest(context)
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil{
		logger.ErrorLogger.Print("Unable to delete, Requested event does not exist.")
		utility.SendInternalServerError(context)
		return
	} 
	err = event.Delete()
	if err != nil {
		utility.SendInternalServerError(context)
		return 
	}
	logger.WarningLogger.Print("Event deleted successfully.")
	// context.JSON(http.StatusOK, utility.GetPrompt("Event deleted successfully."))
	utility.SendSuccessResponse(context, 200)

}