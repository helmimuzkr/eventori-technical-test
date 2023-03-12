package handler

import (
	"eventori/internal/catalogue"
	"eventori/pkg/helper"

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

	catalogueResp := []CatalogueResponse{}
	helper.CopyStruct(&catalogueResp, &catalogueCores)

	return ctx.JSON(helper.SuccessResponse(200, "Get list of models successfully", "success", catalogueResp))
}

func (ch *catalogueHandler) Create(ctx echo.Context) error {
	return nil
}

func (ch *catalogueHandler) GetByID(ctx echo.Context) error {
	return nil
}

func (ch *catalogueHandler) Update(ctx echo.Context) error {
	return nil
}

func (ch *catalogueHandler) Delete(ctx echo.Context) error {
	return nil
}
