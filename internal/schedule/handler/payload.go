package handler

type scheduleResponse struct {
	ID           int    `json:"id"`
	ModelID      int    `json:"model_id"`
	ScheduleDate string `json:"schedule_date"`
}

type scheduleRequest struct {
	ModelID      int    `json:"model_id"`
	ScheduleDate string `json:"schedule_date"`
}
