package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// sample represents data about a record sample.
type sample struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Title string `json:"title"`
}

// samples slice to seed record sample data.
var samples = []sample{
	{Name: "Test1", ID: "1", Title: "Blue Train"},
	{Name: "Test2", ID: "2", Title: "Jeru"},
	{Name: "Test3", ID: "3", Title: "Sarah Vaughan and Clifford Brown"},
}

func main() {
	router := gin.Default()
	router.GET("/samples", getSamples)
	router.GET("/samples/:name", getSampleByName)
	router.POST("/samples", postSamples)
	router.PUT("/samples/:name", putSampleByName)
	router.DELETE("/samples/:name", deleteSampleByName)

	router.Run(":8080")
}

// @Summary Get all Samples
// @Description Get all Samples
// @ID get-samples
// @Produce json
// @Success 200 {object} sample
// @Router /samples [get]
func getSamples(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, samples)
}

// @Summary Create a new Sample
// @Description Add a new sample to the list
// @ID post-sample
// @Accept json
// @Produce json
// @Param sample body sample true "Sample to add"
// @Success 201 {object} sample
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /samples [post]
func postSamples(c *gin.Context) {
	var newsample sample

	if err := c.BindJSON(&newsample); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	// Check if the sample already exists
	for _, s := range samples {
		if s.Name == newsample.Name {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "sample with this name already exists"})
			return
		}
	}

	samples = append(samples, newsample)
	c.IndentedJSON(http.StatusCreated, newsample)
}

// @Summary Get a Sample by Name
// @Description Get a single sample by its name
// @ID get-sample-by-name
// @Produce json
// @Param name path string true "Sample Name"
// @Success 200 {object} sample
// @Failure 404 {object} map[string]string
// @Router /samples/{name} [get]
func getSampleByName(c *gin.Context) {
	name := c.Param("name")

	// Loop through the list of samples, looking for
	// a sample whose name value matches the parameter.
	for _, a := range samples {
		if a.Name == name {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sample not found"})
}

// @Summary Update a Sample by Name
// @Description Update an existing sample
// @ID put-sample-by-name
// @Accept json
// @Produce json
// @Param name path string true "Sample Name"
// @Param sample body sample true "Updated Sample"
// @Success 200 {object} sample
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /samples/{name} [put]
func putSampleByName(c *gin.Context) {
	name := c.Param("name")
	var updatedSample sample

	if err := c.BindJSON(&updatedSample); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	for i, a := range samples {
		if a.Name == name {
			samples[i] = updatedSample
			c.IndentedJSON(http.StatusOK, updatedSample)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sample not found"})
}

// @Summary Delete a Sample by Name
// @Description Delete a sample by its name
// @ID delete-sample-by-name
// @Produce json
// @Param name path string true "Sample Name"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /samples/{name} [delete]
func deleteSampleByName(c *gin.Context) {
	name := c.Param("name")

	// Loop through the list of samples, looking for
	// a sample whose name value matches the parameter.
	for i, a := range samples {
		if a.Name == name {
			// Remove the sample from the slice.
			samples = append(samples[:i], samples[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "sample deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sample not found"})
}
