package base

import (
	"Go-Rampup/config"
	"Go-Rampup/routing"
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (app *App) Initialize() {
	config := config.GetConfig()
	gormDB, err := gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{})
	if err != nil {
		log.Println("Could not connect database", err)
	}
	app.DB = gormDB
	app.Router = routing.GetRouter(gormDB)

}

func (app *App) Server(address string) {
	log.Println("Starting server on: ", address)
	app.Router.Run(address)
}
