package handler

import (
	"eventori/internal/schedule"
	"eventori/pkg/helper"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type scheduleHandler struct {
	srv schedule.Service
}

func New(s schedule.Service) schedule.Handler {
	return &scheduleHandler{srv: s}
}

func (sh *scheduleHandler) GetByModelID(ctx echo.Context) error {
	strID := ctx.Param("model_id")
	modelID, err := strconv.Atoi(strID)
	if err != nil {
		log.Println(err)
		err = helper.SetError(500, "internal server error")
		return ctx.JSON(helper.ErrorResponse(err))
	}

	scheduleCores, err := sh.srv.GetByModelID(modelID)
	if err != nil {
		return ctx.JSON(helper.ErrorResponse(err))
	}

	scheduleResp := []scheduleResponse{}
	helper.CopyStruct(&scheduleResp, &scheduleCores)

	return ctx.JSON(helper.SuccessResponse(200, "Get list of schedules by model id successfully", scheduleResp))
}

func (sh *scheduleHandler) Create(ctx echo.Context) error {
	scheduleReq := scheduleRequest{}
	if err := ctx.Bind(&scheduleReq); err != nil {
		log.Println(err)
		err = helper.SetError(500, "internal server error")
		return ctx.JSON(helper.ErrorResponse(err))
	}

	scheduleCore := schedule.Core{}
	helper.CopyStruct(&scheduleCore, &scheduleReq)

	if err := sh.srv.Create(scheduleCore); err != nil {
		return ctx.JSON(helper.ErrorResponse(err))
	}

	return ctx.JSON(helper.SuccessResponse(201, "Schedule created successfully"))
}

func (sh *scheduleHandler) Update(ctx echo.Context) error {
	scheduleReq := scheduleRequest{}
	if err := ctx.Bind(&scheduleReq); err != nil {
		log.Println(err)
		err = helper.SetError(500, "internal server error")
		return ctx.JSON(helper.ErrorResponse(err))
	}
	strID := ctx.Param("schedule_id")
	scheduleID, err := strconv.Atoi(strID)
	if err != nil {
		log.Println(err)
		err = helper.SetError(500, "internal server error")
		return ctx.JSON(helper.ErrorResponse(err))
	}

	scheduleCore := schedule.Core{}
	helper.CopyStruct(&scheduleCore, &scheduleReq)

	if err := sh.srv.Update(scheduleID, scheduleCore); err != nil {
		return ctx.JSON(helper.ErrorResponse(err))
	}

	return ctx.JSON(helper.SuccessResponse(200, "Schedule updated successfully"))
}

func (sh *scheduleHandler) Delete(ctx echo.Context) error {
	strID := ctx.Param("schedule_id")
	scheduleID, err := strconv.Atoi(strID)
	if err != nil {
		log.Println(err)
		err = helper.SetError(500, "internal server error")
		return ctx.JSON(helper.ErrorResponse(err))
	}

	if err := sh.srv.Delete(scheduleID); err != nil {
		return ctx.JSON(helper.ErrorResponse(err))
	}

	return ctx.JSON(helper.SuccessResponse(200, "Schedule deleted successfully"))
}
