package routes

import (
	"api-gin/controllers"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/age-rating-categories", controllers.GetAllRating)
	r.POST("/age-rating-categories", controllers.CreateRating)
	r.GET("/age-rating-categories/:id", controllers.GetRatingById)
	r.GET("/age-rating-categories/:id/movies", controllers.GetMoviesByAgeRatingCategoryId)
	r.PATCH("/age-rating-categories/:id", controllers.UpdateRating)
	r.DELETE("/age-rating-categories/:id", controllers.DeleteRating)

	r.GET("/movies", controllers.GetAllMovie)
	r.POST("/movies", controllers.CreateMovie)
	r.GET("/movies/:id", controllers.GetMovieById)
	r.PATCH("/movies/:id", controllers.UpdateMovie)
	r.DELETE("/movies/:id", controllers.DeleteMovie)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
