package catalogue

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          int
	ModelName   string
	ImageURL    string
	Description string

	Schedules []Schedule
}

type Schedule struct {
	ID           int
	ModelID      int
	ScheduleDate string
}

type Handler interface {
	GetList(ctx echo.Context) error
	Create(ctx echo.Context) error
	GetByID(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type Service interface {
	GetList() ([]Core, error)
	Create(catalogueCore Core, image *multipart.FileHeader) error
	GetByID(modelID int) (Core, error)
	Update(modelID int, catalogueCore Core, image *multipart.FileHeader) error
	Delete(modelID int) error
}

type Repository interface {
	GetList() ([]Core, error)
	Create(catalogueCore Core) error
	GetByID(modelID int) (Core, error)
	Update(modelID int, catalogueCore Core) error
	Delete(modelID int) error
}
