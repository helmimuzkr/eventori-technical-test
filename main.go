package main

import (
	"eventori/config"
	_catalogueHandler "eventori/internal/catalogue/handler"
	_catalogueRepo "eventori/internal/catalogue/repository"
	_catalogueService "eventori/internal/catalogue/service"
	_scheduleHandler "eventori/internal/schedule/handler"
	_scheduleRepo "eventori/internal/schedule/repository"
	_scheduleService "eventori/internal/schedule/service"
	"eventori/pkg/cloud"
	"eventori/pkg/database"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := config.InitConfig()

	db := database.InitMysql(conf)
	cloudinary := cloud.NewCloudinary(conf)

	catalogueRepo := _catalogueRepo.New(db)
	catalogueSrv := _catalogueService.New(catalogueRepo, cloudinary.Upload)
	catalogueHandler := _catalogueHandler.New(catalogueSrv)

	scheduleRepo := _scheduleRepo.New(db)
	scheduleSrv := _scheduleService.New(scheduleRepo)
	scheduleHandler := _scheduleHandler.New(scheduleSrv)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom}, method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	g := e.Group("/models")
	g.GET("/list", catalogueHandler.GetList)
	g.POST("/create", catalogueHandler.Create)
	g.GET("/:model_id", catalogueHandler.GetByID)
	g.DELETE("/:model_id", catalogueHandler.Delete)
	g.PATCH("/update/:model_id", catalogueHandler.Update)

	g.GET("/schedules/:model_id", scheduleHandler.GetByModelID)
	g.POST("/schedules/create", scheduleHandler.Create)
	g.PATCH("/schedules/update/:schedule_id", scheduleHandler.Update)
	g.DELETE("/schedules/:schedule_id", scheduleHandler.Delete)

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
