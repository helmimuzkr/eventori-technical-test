package repository

import (
	"eventori/internal/schedule"
	"log"
	"time"
)

type ModelSchedule struct {
	ID           int        `gorm:"column:Id"`
	ModelID      int        `gorm:"column:Model_id"`
	ScheduleDate *time.Time `gorm:"column:Schedule_date"`
	CreatedAt    *time.Time `gorm:"column:Created_at"`
	UpdatedAt    *time.Time `gorm:"column:Updated_at"`
	DeletedAt    *time.Time `gorm:"column:Deleted_at"`
}

func ToModel(c schedule.Core) ModelSchedule {
	m := ModelSchedule{}
	m.ID = c.ID
	m.ModelID = c.ModelID

	if c.ScheduleDate != "" {
		converted, err := time.Parse("2006-01-02", c.ScheduleDate)
		if err != nil {
			log.Println("error parsing time in model schedule service")
		}
		m.ScheduleDate = &converted
	}

	return m
}

func ToCores(m ...ModelSchedule) []schedule.Core {
	cores := []schedule.Core{}
	for _, v := range m {
		c := schedule.Core{}
		c.ID = v.ID
		c.ModelID = v.ModelID
		c.ScheduleDate = v.ScheduleDate.Format("2006-01-02")
		cores = append(cores, c)
	}

	return cores
}
