package schedule

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           int
	ModelID      int
	ScheduleDate string
}

type Handler interface {
	GetByModelID(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type Service interface {
	GetByModelID(modelID int) ([]Core, error)
	Create(scheduleCore Core) error
	Update(scheduleID int, shceduleCore Core) error
	Delete(scheduleID int) error
}

type Repository interface {
	GetByModelID(modelID int) ([]Core, error)
	Create(scheduleCore Core) error
	Update(scheduleID int, shceduleCore Core) error
	Delete(scheduleID int) error
}
