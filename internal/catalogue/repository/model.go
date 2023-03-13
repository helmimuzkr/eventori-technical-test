package repository

import (
	"eventori/internal/catalogue"
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
	ScheduleDate time.Time  `gorm:"column:Schedule_date"`
	CreatedAt    *time.Time `gorm:"column:Created_at"`
	UpdatedAt    *time.Time `gorm:"column:Updated_at"`
	DeletedAt    *time.Time `gorm:"column:Deleted_at"`
}

func ToModel(c catalogue.Core) ModelCatalogue {
	m := ModelCatalogue{}
	m.ModelName = c.ModelName
	m.ImageURL = c.ImageURL
	m.Description = c.Description

	return m
}

func ToCores(m ...ModelCatalogue) []catalogue.Core {
	cores := []catalogue.Core{}
	for _, v := range m {
		core := catalogue.Core{}
		core.ID = v.ID
		core.ModelName = v.ModelName
		core.ImageURL = v.ImageURL
		core.Description = v.Description

		scheduleCores := []catalogue.Schedule{}
		for i := range v.Schedules {
			c := catalogue.Schedule{}
			c.ID = v.Schedules[i].ID
			c.ModelID = v.Schedules[i].ModelID
			c.ScheduleDate = v.Schedules[i].ScheduleDate.Format("2006-01-02")
			scheduleCores = append(scheduleCores, c)
		}
		core.Schedules = scheduleCores

		cores = append(cores, core)
	}

	return cores
}
