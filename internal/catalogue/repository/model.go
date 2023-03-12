package repository

import (
	"time"
)

type ModelCatalogue struct {
	ID          int        `gorm:"column:Id"`
	ModelName   string     `gorm:"column:Model_name"`
	ImageURL    string     `gorm:"column:Image_url"`
	Description string     `gorm:"column:Description"`
	CreatedAt   *time.Time `gorm:"column:Created_at"`
	UpdatedAt   *time.Time `gorm:"column:Updated_at"`
	DeletedAt   *time.Time `gorm:"column:Deleted_at"`

	Schedules []ModelSchedule `gorm:"foreignKey:Model_id"`
}

type ModelSchedule struct {
	ID           int        `gorm:"column:Id"`
	ModelID      int        `gorm:"column:Model_id"`
	ScheduleDate *time.Time `gorm:"column:Schedule_date"`
	CreatedAt    *time.Time `gorm:"column:Created_at"`
	UpdatedAt    *time.Time `gorm:"column:Updated_at"`
	DeletedAt    *time.Time `gorm:"column:Deleted_at"`
}
