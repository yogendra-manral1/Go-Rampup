package routing

import (
	"github.com/gin-contrib/cors"
	"Go-Rampup/apps/auth"
	"Go-Rampup/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()

	router.Use(middlewares.PanicMiddleware())

	// cors middleware allows CORS request, we have configured it based on the application.
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}

	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AllowCredentials = true
	// Application might receive some headers in the request data, so the cors middleware must know about them.
	corsConfig.AddAllowMethods([]string{"PATCH", "OPTIONS", "DELETE"}...)
	router.Use(cors.New(corsConfig))

	apiRouter := router.Group("/api/v1")

	// Add api routes
	userAuthCtrl := auth.UserAuthController{DB: db}
	userAuthCtrl.SetRoutes(apiRouter.Group("/auth"))
	return router
}
