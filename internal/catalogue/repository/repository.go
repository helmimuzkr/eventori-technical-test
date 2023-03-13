package repository

import (
	"errors"
	"eventori/internal/catalogue"
	"fmt"

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

	return ToCores(catalogueModels...), nil
}

func (cr *catalogueRepository) Create(catalogueCore catalogue.Core) error {
	catalogueModel := ToModel(catalogueCore)
	tx := cr.db.Create(&catalogueModel)
	if tx.Error != nil {
		return tx.Error
	}
	
	return nil
}

func (cr *catalogueRepository) GetByID(modelID int) (catalogue.Core, error) {
	catalogueModel := ModelCatalogue{}
	tx := cr.db.Preload("Schedules").First(&catalogueModel, modelID)
	if tx.Error != nil {
		return catalogue.Core{}, tx.Error
	}
	
	return ToCores(catalogueModel)[0], nil
}

func (cr *catalogueRepository) Update(modelID int, catalogueCore catalogue.Core) error {
	catalogueModel := ToModel(catalogueCore)
	tx := cr.db.Where("id = ?", modelID).Updates(&catalogueModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed because no rows affected")
	}

	return nil
}

func (cr *catalogueRepository) Delete(modelID int) error {
	tx := cr.db.Delete(ModelCatalogue{}, modelID)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed because no rows affected")
	}

	return nil
}
