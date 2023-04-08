package controllers

import (
	"api-gin/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AgeRatingCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Get All Rating godoc
// @Summary Get All Age Rating Category
// @Description Get List of Age Rating Category
// @Tags AgeRatingCategory
// @Produce json
// @Success 200 {object} []models.AgeRatingCategory
// @Router /age-rating-categories [get]
func GetAllRating(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var ratings []models.AgeRatingCategory

	db.Find(&ratings)

	c.JSON(http.StatusOK, gin.H{"data": ratings})

}

// Create a Rating godoc
// @Summary Create Age Rating Category
// @Description create new Age Rating Category
// @Tags AgeRatingCategory
// @Param Body body AgeRatingCategoryInput true "the body to create new age rating category"
// @Produce json
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories [post]
func CreateRating(c *gin.Context) {
	var input AgeRatingCategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating := models.AgeRatingCategory{Name: input.Name, Description: input.Description}
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	db.Create(&rating)
	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// Get a Rating godoc
// @Summary Get an Age Rating Category by id
// @Description Get one Age Rating Category by id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Eating Category Id"
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories/{id} [get]
func GetRatingById(c *gin.Context) {
	var rating models.AgeRatingCategory
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// Get movies from one Rating godoc
// @Summary Get movies by Age Rating Category by id
// @Description Get all movies of Age Rating Category by id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Eating Category Id"
// @Success 200 {object} []models.Movie
// @Router /age-rating-categories/{id}/movies [get]
func GetMoviesByAgeRatingCategoryId(c *gin.Context) {
	var movies []models.Movie
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("age_rating_category_id = ?", c.Param("id")).Find(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movies})

}

// update rating godoc
// @Summary update an Age Rating Category by id
// @Description update one Age Rating Category by id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Eating Category Id"
// @Param Body body AgeRatingCategoryInput true "the body to create new age rating category"
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories/{id} [patch]
func UpdateRating(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// get rating if exist
	var rating models.AgeRatingCategory
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	var input AgeRatingCategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputRating models.AgeRatingCategory

	updatedInputRating.Name = input.Name
	updatedInputRating.Description = input.Description
	updatedInputRating.UpdatedAt = time.Now()

	db.Model(&rating).Updates(updatedInputRating)
	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// delete a Rating godoc
// @Summary delete an Age Rating Category by id
// @Description delete one Age Rating Category by id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Eating Category Id"
// @Success 200 {object} map[string]boolean
// @Router /age-rating-categories/{id} [delete]
func DeleteRating(c *gin.Context) {
	var rating models.AgeRatingCategory
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	db.Delete(&rating)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
