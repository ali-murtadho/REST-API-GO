package controllers

import (
	"api-gin/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieInput struct {
	Title               string `json:"title"`
	Year                int    `json:"year"`
	AgeRatingCategoryID int    `json:"age_rating_category_id"`
}

// Get All Rating godoc
// @Summary Get All Movie
// @Description Get List of Movie
// @Tags Movie
// @Produce json
// @Success 200 {object} []models.Movie
// @Router /movies [get]
func GetAllMovie(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var movies []models.Movie

	db.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})

}

// Create a Rating godoc
// @Summary Create Movie
// @Description create new Movie
// @Tags Movie
// @Param Body body MovieInput true "the body to create new Movie"
// @Produce json
// @Success 200 {object} models.Movie
// @Router /movies [post]
func CreateMovie(c *gin.Context) {
	var input MovieInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie := models.Movie{Title: input.Title, Year: input.Year, AgeRatingCategoryID: input.AgeRatingCategoryID}
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	db.Create(&movie)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// Get a Rating godoc
// @Summary Get an Movie by id
// @Description Get one Movie by id
// @Tags Movie
// @Produce json
// @Param id path string true "Movie Id"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [get]
func GetMovieById(c *gin.Context) {
	var movie models.Movie
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// update rating godoc
// @Summary update an Movie by id
// @Description update one Movie by id
// @Tags Movie
// @Produce json
// @Param id path string true "Movie Id"
// @Param Body body MovieInput true "the body to create new Movie"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [patch]
func UpdateMovie(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// get rating if exist
	var movie models.Movie
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	var input MovieInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updateInput models.Movie

	updateInput.Title = input.Title
	updateInput.Year = input.Year
	updateInput.AgeRatingCategoryID = input.AgeRatingCategoryID
	updateInput.UpdatedAt = time.Now()

	db.Model(&movie).Updates(updateInput)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// delete a Rating godoc
// @Summary delete an Movie by id
// @Description delete one Movie by id
// @Tags Movie
// @Produce json
// @Param id path string true "Movie Id"
// @Success 200 {object} map[string]boolean
// @Router /movies/{id} [delete]
func DeleteMovie(c *gin.Context) {
	var movie models.Movie
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	db.Delete(&movie)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
