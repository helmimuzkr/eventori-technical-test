package handler

type catalogueResponse struct {
	ID          int    `json:"id"`
	ModelName   string `json:"model_name"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`

	Schedules []Schedule `json:"schedules"`
}

type Schedule struct {
	ID           int    `json:"id"`
	ModelID      int    `json:"model_id"`
	ScheduleDate string `json:"schedule_date"`
}

type catalogueRequest struct {
	ModelName   string `json:"model_name" form:"model_name"`
	Description string `json:"description" form:"description"`
}
