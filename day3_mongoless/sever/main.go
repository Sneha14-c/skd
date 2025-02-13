package main

import (
    "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Flight struct {
	Id string
	Number string
	AirlineName string
	Source string
	Destination string
	Capacity int 
	Price float32 
}
func readAllFlights(c *gin.Context) {
	flights := []Flight{
		{Id: "120", Number: "AI 845",
			AirlineName: "Air India",
			Source:      "Mumbai", Destination: "Abu Dhabi", Capacity: 180,
			Price: 15000.00},
		{Id: "125", Number: "AI 945",
			AirlineName: "Air India",
			Source:      "Delhi", Destination: "Abu Dhabi", Capacity: 180,
			Price: 16000.00},
	}
	fmt.Println(flights)

	c.JSON(http.StatusOK,flights)
}
func readFlightById(c *gin.Context){
	id := c.Param("id")
	flight := Flight {Id:id,Number:"AI 845",AirlineName:"Air India",Source:"Mumbai",Destination:"Abu Dhabi",Capacity:180,Price:15000.00}
	

	fmt.Println(flight)

	c.JSON(http.StatusOK,flight)
}
func createFlight(c *gin.Context){
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err !=nil{
		c.JSON(http.StatusInternalServerError,
		gin.H{"error" : "Server Error ." + err.Error()})
		return
	}
	createdFlight := Flight {Id:"1001",Number:"AI 845",AirlineName:"Air India",Source:"Mumbai",Destination:"Abu Dhabi",Capacity:180,Price:15000.00}
	c.JSON(http.StatusCreated,
	gin.H{"message" : "Flight created successfully.",
"flight" : createdFlight})
	
}


func updateFlight(c *gin.Context){
	id := c.Param("id")
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err !=nil{
		c.JSON(http.StatusInternalServerError,
		gin.H{"error" : "Server Error ." + err.Error()})
		return
	}
	updatedFlight := Flight {Id:id,Number:"AI 845",AirlineName:"Air India",Source:"Mumbai",Destination:"Abu Dhabi",Capacity:180,Price:20000.00}
	c.JSON(http.StatusOK,
	gin.H{"message" : "Flight price is updated successfully.",
"flight" : updatedFlight})
	
}

func deleteFlight(c *gin.Context){
	//id := c.Param("id")
	
	c.JSON(http.StatusOK,
	gin.H{"message" : "Flight deleted successfully."})
	
}




func main() {
	//router
	r :=gin.Default()
//routes | API endpoints
    r.GET("/flights",readAllFlights)
	r.GET("/flights/:id",readFlightById)
	r.POST("/flights",createFlight)
	r.PUT("/flights/:id",updateFlight)
	r.DELETE("/flights/:id",deleteFlight)
	//server (default portal 8080)
	r.Run(":8080")

}
