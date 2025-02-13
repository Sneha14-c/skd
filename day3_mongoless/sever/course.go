package main

import (
    "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
type Course struct {
	Id string
	Name string
	Duration string
	Triner_name string
	Level string
	Description string
	Price float32 
}
func readAllCourses(c *gin.Context) {
	Courses := []Course{
		{Id: "120", Name: "sneha",
			Duration: "3 months",
			Triner_name: "Nithin", Level: "Beginner", Description:"python for data science",
			Price: 4000.00},
			{Id: "150", Name: "sushmitha",
			Duration: "3 months",
			Triner_name: "Mahesh", Level: "Advance", Description:"DEART language",
			Price: 6000.00},
	}
	fmt.Println(Courses)

	c.JSON(http.StatusOK,Courses)
}

func updateCourse(c *gin.Context){
	id := c.Param("id")
	var jbodyCourse Course
	err := c.BindJSON(&jbodyCourse)
	if err !=nil{
		c.JSON(http.StatusInternalServerError,
		gin.H{"error" : "Server Error ." + err.Error()})
		return
	}
	updatedCourse := Course {Id: id, Name: "sneha",
	Duration: "3 months",
	Triner_name: "Nithin", Level: "Beginner", Description:"python for data science",
	Price: 5000.00}
	c.JSON(http.StatusOK,
	gin.H{"message" : "course price is updated successfully.",
"flight" : updatedCourse})
	
}

func main() {
	//router
	r :=gin.Default()
//routes | API endpoints
    r.GET("/Courses",readAllCourses)
	// r.GET("/flights/:id",readFlightById)
	// r.POST("/flights",createFlight)
	 r.PUT("/flights/:id",updateCourse)
	// r.DELETE("/flights/:id",deleteFlight)
	//server (default portal 8080)
	r.Run(":8080")

}