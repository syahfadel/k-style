package routers

import (
	"unnispick/controllers"
	"unnispick/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func StartService(db *gorm.DB) *echo.Echo {

	unnispickService := services.UnnispickService{
		DB: db,
	}

	unnispickController := controllers.UnnispickController{
		DB:               db,
		UnnispickService: unnispickService,
	}

	app := echo.New()
	app.POST("/member", unnispickController.CreateMember)
	app.PUT("/member/:id", unnispickController.UpdateMember)
	app.DELETE("/member/:id", unnispickController.DeleteMember)
	app.GET("/member", unnispickController.GetAllMember)
	app.GET("/product/:id", unnispickController.GetProductById)
	app.POST("/like", unnispickController.Like)
	return app
}
