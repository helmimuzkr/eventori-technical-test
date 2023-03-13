package service

import (
	"eventori/internal/schedule"
	"eventori/pkg/helper"
	"log"
)

type scheduleService struct {
	repo schedule.Repository
}

func New(r schedule.Repository) schedule.Service {
	return &scheduleService{repo: r}
}

func (ss *scheduleService) GetByModelID(modelID int) ([]schedule.Core, error) {
	scheduleCores, err := ss.repo.GetByModelID(modelID)
	if err != nil {
		log.Println(err)
		return nil, helper.SetError(500, "internal server error")
	}

	return scheduleCores, nil
}

func (ss *scheduleService) Create(scheduleCore schedule.Core) error {
	if scheduleCore.ScheduleDate == "" {
		log.Println("schedule date is required")
		return helper.SetError(400, "bad request")
	}

	if err := ss.repo.Create(scheduleCore); err != nil {
		log.Println(err)
		return helper.SetError(500, "internal server error")
	}

	return nil
}

func (ss *scheduleService) Update(scheduleID int, scheduleCore schedule.Core) error {
	if err := ss.repo.Update(scheduleID, scheduleCore); err != nil {
		log.Println(err)
		return helper.SetError(500, "internal server error")
	}

	return nil
}

func (ss *scheduleService) Delete(scheduleID int) error {
	if err := ss.repo.Delete(scheduleID); err != nil {
		log.Println(err)
		return helper.SetError(500, "internal server error")
	}

	return nil
}
