package handler

import "time"

type CatalogueResponse struct {
	ID          int    `json:"id"`
	ModelName   string `json:"model_name"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`

	Schedules []Schedule `json:"schedules"`
}

type Schedule struct {
	ID           int        `json:"id"`
	ModelID      int        `json:"model_id"`
	ScheduleDate *time.Time `json:"schedule_date"`
}
