package routes

import (
	"crud_project/models"
	"crud_project/utility"
	"net/http"
	"strconv"
	"crud_project/logger"
	"github.com/gin-gonic/gin"
)


func GetEvents(context *gin.Context){
	// function to get all the events from the database.
	event, err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError, utility.GetPrompt("could not fetch evemts. try again later."))
		return
	}
	
	data := utility.GetPrompt("sucessfully fetched events", event) 
	context.JSON(http.StatusOK, data) 
}

func CreateEvent(context *gin.Context) {
	// Function to create event, post request for event creation.

	var event models.Event  // Variable of type models.Event

	err := context.ShouldBindJSON(&event) // to read data from the request, and map it to the event struct

	logger.InfoLogger.Println("Working with create event function")

	if err!= nil{
		logger.ErrorLogger.Print("Error while reading the input")
		data := utility.GetPrompt("cannot parse the request")
		context.JSON(http.StatusBadRequest, data)
		return 
	}

	err = event.Save() 

	if err != nil{
		logger.ErrorLogger.Print("Error in event creation")
		data := utility.GetPrompt("could not create events")
		context.JSON(http.StatusInternalServerError, data)
		return
	}

	logger.InfoLogger.Println("Event created successfully")

    data := utility.GetPrompt("Event Created!", event)
	context.JSON(http.StatusCreated, data)
}

func GetEvent(context *gin.Context) { 
	// gets single event by id

	eventid, err := strconv.ParseInt(context.Param("id"),10,64) // fetch id from path
	if err != nil{
		logger.ErrorLogger.Print("Unable to parse request")
		context.JSON(http.StatusBadRequest, utility.GetPrompt("Could not parse request data."))
		return
	}

	event, err := models.GetEventByID(eventid)
	if err != nil{
		logger.ErrorLogger.Print("Unable to fetch event")
		context.JSON(http.StatusInternalServerError, utility.GetPrompt("Could not fetch event."))
		return
	} 

	context.JSON(http.StatusOK, utility.GetPrompt("sucessfully fetched event", event))
}

func UpdateEvent(context *gin.Context) {
	// Function to update event details
	eventid, err := strconv.ParseInt(context.Param("id"),10,64) // fetch id from path
	if err != nil{
		logger.ErrorLogger.Print("Unable to parse request")
		context.JSON(http.StatusBadRequest, utility.GetPrompt("Could not parse request data."))
		return
	}
	_, err = models.GetEventByID(eventid) // check if event with the given id is there or not
	data := utility.GetPrompt("event does not exist")
	if err != nil{
		context.JSON(http.StatusInternalServerError, data)
		return
	}

	var updatedEvent models.Event    // new variable to store data send in the request by user
	err = context.ShouldBindJSON(&updatedEvent)   // extracting data from request and storing in updatedEvent
	
	if err!=nil{
		context.JSON(http.StatusBadRequest, utility.GetPrompt("Could not update event"))
		return
	}
	updatedEvent.ID = eventid // copying id in the current event, so that new id is not created
	err = updatedEvent.Update()  // calling the update method to push data to db.

	if err !=nil {
		context.JSON(http.StatusInternalServerError, utility.GetPrompt("Could not update event"))
		return
	}
	context.JSON(http.StatusOK,utility.GetPrompt("event updated sucessfully"))
}

func DeleteEvent(context *gin.Context)  {
	// Function to delete event.
	eventId, err := strconv.ParseInt(context.Param("id"),10, 64)

	if err != nil{
		logger.ErrorLogger.Print("Unable to parse request")
		context.JSON(http.StatusBadRequest, utility.GetPrompt("Could not parse request data."))
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil{
		logger.ErrorLogger.Print("Unable to delete, Requested event does not exist.")
		context.JSON(http.StatusInternalServerError, utility.GetPrompt("Requested event does not exist."))
		return
	} 
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, utility.GetPrompt("Could not delete event"))
		return 
	}
	logger.WarningLogger.Print("Event deleted successfully.")
	context.JSON(http.StatusOK, utility.GetPrompt("Event deleted successfully."))

}