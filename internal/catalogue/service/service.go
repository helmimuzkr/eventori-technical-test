package service

import (
	"eventori/internal/catalogue"
	"eventori/pkg/helper"
)

type catalogueService struct {
	repo catalogue.Repository
}

func New(r catalogue.Repository) catalogue.Service {
	return &catalogueService{repo: r}
}

func (cs *catalogueService) GetList() ([]catalogue.Core, error) {
	catalogueCores, err := cs.repo.GetList()
	if err != nil {
		return nil, helper.SetError(500, "internal server error")
	}

	return catalogueCores, nil
}
func (cs *catalogueService) Create(catalogueCore catalogue.Core) error {
	return nil
}
func (cs *catalogueService) GetByID(modelID int) (catalogue.Core, error) {
	return catalogue.Core{}, nil
}
func (cs *catalogueService) Update(modelID int) error {
	return nil
}
func (cs *catalogueService) Delete(modelID int) error {
	return nil
}
