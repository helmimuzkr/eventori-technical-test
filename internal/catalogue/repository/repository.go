package repository

import (
	"eventori/internal/catalogue"
	"eventori/pkg/helper"

	"gorm.io/gorm"
)

type catalogueRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) catalogue.Repository {
	return &catalogueRepository{db: db}
}

func (cr *catalogueRepository) GetList() ([]catalogue.Core, error) {
	catalogueModels := []ModelCatalogue{}
	tx := cr.db.Preload("Schedules").Find(&catalogueModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	catalogueCores := []catalogue.Core{}
	helper.CopyStruct(&catalogueCores, &catalogueModels)

	return catalogueCores, nil
}

func (cr *catalogueRepository) Create(catalogueCore catalogue.Core) error {
	return nil
}
func (cr *catalogueRepository) GetByID(modelID int) (catalogue.Core, error) {
	return catalogue.Core{}, nil
}
func (cr *catalogueRepository) Update(modelID int) error {
	return nil
}
func (cr *catalogueRepository) Delete(modelID int) error {
	return nil
}
