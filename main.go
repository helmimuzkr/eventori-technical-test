package main

import (
	"eventori/config"
	_catalogueHandler "eventori/internal/catalogue/handler"
	_catalogueRepo "eventori/internal/catalogue/repository"
	_catalogueService "eventori/internal/catalogue/service"
	"eventori/pkg/database"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := database.InitMysql(conf)

	catalogueRepo := _catalogueRepo.New(db)
	catalogueSrv := _catalogueService.New(catalogueRepo)
	catalogueHandler := _catalogueHandler.New(catalogueSrv)

	e := echo.New()
	g := e.Group("/models")

	g.GET("/list", catalogueHandler.GetList)

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
