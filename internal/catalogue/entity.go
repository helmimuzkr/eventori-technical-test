package catalogue

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          int
	ModelName   string
	ImageURL    string
	Description string

	Schedules []struct {
		ID           int
		ModelID      int
		ScheduleDate *time.Time
	}
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
	Create(catalogueCore Core) error
	GetByID(modelID int) (Core, error)
	Update(modelID int) error
	Delete(modelID int) error
}

type Repository interface {
	GetList() ([]Core, error)
	Create(catalogueCore Core) error
	GetByID(modelID int) (Core, error)
	Update(modelID int) error
	Delete(modelID int) error
}
