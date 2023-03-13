package handler

import (
	"eventori/internal/catalogue"
	"eventori/pkg/helper"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type catalogueHandler struct {
	srv catalogue.Service
}

func New(s catalogue.Service) catalogue.Handler {
	return &catalogueHandler{srv: s}
}

func (ch *catalogueHandler) GetList(ctx echo.Context) error {
	catalogueCores, err := ch.srv.GetList()
	if err != nil {
		return ctx.JSON(helper.ErrorResponse(err))
	}

	catalogueResp := []catalogueResponse{}
	helper.CopyStruct(&catalogueResp, &catalogueCores)

	return ctx.JSON(helper.SuccessResponse(200, "Get list of models successfully", catalogueResp))
}

func (ch *catalogueHandler) Create(ctx echo.Context) error {
	catalogueReq := catalogueRequest{}
	if err := ctx.Bind(&catalogueReq); err != nil {
		log.Println(err)
		err = helper.SetError(500, "internal server error")
		return ctx.JSON(helper.ErrorResponse(err))
	}
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		log.Println(err)
		err = helper.SetError(500, "internal server error")
		return ctx.JSON(helper.ErrorResponse(err))
	}

	catalogueCore := catalogue.Core{}
	helper.CopyStruct(&catalogueCore, &catalogueReq)

	if err := ch.srv.Create(catalogueCore, fileHeader); err != nil {
		return ctx.JSON(helper.ErrorResponse(err))
	}

	return ctx.JSON(helper.SuccessResponse(201, "Model created successfully"))
}

func (ch *catalogueHandler) GetByID(ctx echo.Context) error {
	strID := ctx.Param("model_id")
	modelID, err := strconv.Atoi(strID)
	if err != nil {
		log.Println(err)
		err = helper.SetError(500, "internal server error")
		return ctx.JSON(helper.ErrorResponse(err))
	}

	catalogueCore, err := ch.srv.GetByID(modelID)
	if err != nil {
		return ctx.JSON(helper.ErrorResponse(err))
	}
	catalogueResp := catalogueResponse{}
	helper.CopyStruct(&catalogueResp, &catalogueCore)

	return ctx.JSON(helper.SuccessResponse(200, "Get model successfully", catalogueResp))
}

func (ch *catalogueHandler) Update(ctx echo.Context) error {
	strID := ctx.Param("model_id")
	modelID, err := strconv.Atoi(strID)
	if err != nil {
		log.Println(err)
		err = helper.SetError(500, "internal server error")
		return ctx.JSON(helper.ErrorResponse(err))
	}
	catalogueReq := catalogueRequest{}
	if err := ctx.Bind(&catalogueReq); err != nil {
		log.Println(err)
		err = helper.SetError(500, "internal server error")
		return ctx.JSON(helper.ErrorResponse(err))
	}
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		log.Println(err)
	}

	catalogueCore := catalogue.Core{}
	helper.CopyStruct(&catalogueCore, &catalogueReq)

	if err := ch.srv.Update(modelID, catalogueCore, fileHeader); err != nil {
		return ctx.JSON(helper.ErrorResponse(err))
	}

	return ctx.JSON(helper.SuccessResponse(200, "Model updated successfully"))
}

func (ch *catalogueHandler) Delete(ctx echo.Context) error {
	strID := ctx.Param("model_id")
	modelID, err := strconv.Atoi(strID)
	if err != nil {
		log.Println(err)
		err = helper.SetError(500, "internal server error")
		return ctx.JSON(helper.ErrorResponse(err))
	}

	if err := ch.srv.Delete(modelID); err != nil {
		return ctx.JSON(helper.ErrorResponse(err))
	}

	return ctx.JSON(helper.SuccessResponse(200, "Model deleted successfully"))
}
